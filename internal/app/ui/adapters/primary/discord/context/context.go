package context

import (
	"context"
	"errors"
	"time"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/rs/zerolog"
)

type CommandContext struct {
	log     zerolog.Logger
	user    tempest.Snowflake
	ctx     context.Context
	tempest *tempest.Client
	*tempest.CommandInteraction
}

func getUserId(itx *tempest.CommandInteraction) (tempest.Snowflake, error) {
	if itx.Member != nil {
		if itx.Member.User != nil {
			return itx.Member.User.ID, nil
		}
	}

	if itx.User != nil {
		return itx.User.ID, nil
	}

	return 0, errors.New("Failed to get user ID")
}

func InitContext(client *tempest.Client, itx *tempest.CommandInteraction, lvl zerolog.Level) (*CommandContext, error) {
	ctx := logger.InitContextLogger(context.Background(), lvl)
	log := logger.FromContext(ctx)
	log.Debug().Msg("Initializing command context")

	log.Debug().Msg("Getting user ID")
	user, err := getUserId(itx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user ID")
		return nil, err
	}

	tex := &CommandContext{
		log:                log,
		CommandInteraction: itx,
		user:               user,
		ctx:                ctx,
		tempest:            client,
	}

	log.Debug().Msg("Deferring command interaction")
	if err := itx.Defer(true); err != nil {
		err := errors.New("Failed to defer command interaction")
		tex.SendLinearReply("Command failed with an unknown error!", true)
		return tex, err
	}

	log.Debug().Msg("Returning command context")
	return tex, nil
}

func (t *CommandContext) Logger() zerolog.Logger {
	return t.log
}

func (t *CommandContext) EditResponse(message string, ephemeral bool) error {
	return t.EditReply(tempest.ResponseMessageData{
		Content: message,
	}, ephemeral)
}

func (t *CommandContext) SendRequesterPrivateMessage(msg string) (tempest.Message, error) {
	return t.tempest.SendPrivateMessage(t.user, tempest.Message{
		Content: msg,
	})
}

func (t *CommandContext) Context() context.Context {
	return t.ctx
}

func (t *CommandContext) Deadline() (deadline time.Time, ok bool) {
	return t.ctx.Deadline()
}

func (t *CommandContext) Done() <-chan struct{} {
	return t.ctx.Done()
}

func (t *CommandContext) Err() error {
	return t.ctx.Err()
}

func (t *CommandContext) Value(key interface{}) interface{} {
	return t.ctx.Value(key)
}
