package mem_workflow

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type WorkerOpt func(*Worker)

func WithLogLevel(lvl zerolog.Level) WorkerOpt {
	return func(w *Worker) {
		w.logger = w.logger.Level(lvl)
	}
}

type Worker struct {
	logLevel    zerolog.Level
	logger      zerolog.Logger
	actSvc      service.Activity
	notSvc      service.Notification
	workChan    chan workflow.ExecutionInput
	activeTasks uint32
}

func NewWorker(actSvc service.Activity, notSvc service.Notification, c chan workflow.ExecutionInput, opts ...WorkerOpt) *Worker {
	w := &Worker{
		logLevel:    zerolog.InfoLevel,
		workChan:    c,
		actSvc:      actSvc,
		notSvc:      notSvc,
		activeTasks: 0,
	}

	for _, o := range opts {
		o(w)
	}
	log := logger.InitLogger(w.logLevel, logger.WithField("component", "worker"))
	w.logger = log

	return w
}

func (w *Worker) incrementRunningTasks() {
	atomic.AddUint32(&w.activeTasks, 1)
}

func (w *Worker) decrementRunningTasks() {
	atomic.AddUint32(&w.activeTasks, ^uint32(0))
}

func (w *Worker) RunningTasks() uint32 {
	return atomic.LoadUint32(&w.activeTasks)
}

func (m *Worker) Start(ctx context.Context) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case wrk := <-m.workChan:
				log.Debug().Msg("Received work")
				m.incrementRunningTasks()
				m.runExecution(wrk)
				m.decrementRunningTasks()
			}
		}
	}()
	return nil
}

func (m *Worker) Close(ctx context.Context) error {
	timeout := time.After(120 * time.Second)
	for {
		if m.RunningTasks() == 0 {
			return nil
		}
		select {
		case <-ctx.Done():
			return errors.New("context cancelled")
		case <-timeout:
			return errors.New("timeout waiting for worker to close")
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func (w *Worker) closeExecution(ctx context.Context, ex workflow.ExecutionIdentifier, result workflow.Status, output workflow.ExecutionResult) error {
	return w.actSvc.CloseExecution(ctx, ex, result, output)
}

func (w *Worker) runExecution(input workflow.ExecutionInput) {
	ctx := logger.InitContextLogger(context.Background(), zerolog.DebugLevel) // TODO: make log level dynamic

	log.Debug().Msg("Validating execution id")
	exId, err := workflow.ExecutionIdentifierFromString(input.ExecutionIdentifier())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse execution id")
	}

	var status workflow.Status
	var result workflow.ExecutionResult
	var wf workflow.WorkflowName

	log.Debug().Msg("Running execution")
	switch input := input.(type) {
	case *workflow.ProvisionNodeInput:
		wf = workflow.WorkflowNameProvisionNode
		status, result = w.provisionNodeWorkflow(ctx, input)
	case *workflow.DeprovisionNodeInput:
		wf = workflow.WorkflowNameDeprovisionNode
		status, result = w.deprovisionNodeWorkflow(ctx, input)
	default:
		err = errors.New("unknown execution type")
		log.Error().Err(err).Msg("Unknown execution type")
		return
	}

	log.Debug().Msg("Closing execution")
	if err = w.closeExecution(ctx, exId, status, result); err != nil {
		log.Error().Err(err).Msg("Failed to close execution")
		return
	}

	log.Debug().Msg("Signaling execution complete")
	err = w.notSvc.NotifyExecutionCompletion(ctx, exId, wf, status, result)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to signal execution complete")
	}

	log.Debug().Msg("Execution complete")
}
