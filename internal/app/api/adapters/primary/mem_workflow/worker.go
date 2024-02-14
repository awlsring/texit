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

type Worker struct {
	logger      zerolog.Logger
	actSvc      service.Activity
	workChan    chan workflow.ExecutionInput
	activeTasks uint32
}

func NewWorker(actSvc service.Activity, c chan workflow.ExecutionInput) *Worker {
	log := logger.FromContext(logger.InitContextLogger(context.Background(), zerolog.DebugLevel))
	return &Worker{
		logger:      log,
		workChan:    c,
		actSvc:      actSvc,
		activeTasks: 0,
	}
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

	log.Debug().Msg("Running execution")
	switch input := input.(type) {
	case *workflow.ProvisionNodeInput:
		status, result = w.provisionNodeWorkflow(ctx, input)
	case *workflow.DeprovisionNodeInput:
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
}
