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
	log              zerolog.Logger
	requester        tempest.Snowflake
	requesterRoles   []tempest.Snowflake
	ctx              context.Context
	tempest          *tempest.Client
	isPrivateMessage bool
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

func isPrivateMessage(itx *tempest.CommandInteraction) bool {
	if itx.User != nil {
		return true
	}

	return false
}

func getRoles(itx *tempest.CommandInteraction) []tempest.Snowflake {
	if itx.Member != nil {
		return itx.Member.RoleIDs
	}
	return nil
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
		requester:          user,
		requesterRoles:     getRoles(itx),
		isPrivateMessage:   isPrivateMessage(itx),
		ctx:                ctx,
		tempest:            client,
	}

	log.Debug().Msg("Returning command context")
	return tex, nil
}

func (t *CommandContext) Logger() zerolog.Logger {
	return t.log
}

func (t *CommandContext) Requester() tempest.Snowflake {
	return t.requester
}

func (t *CommandContext) RequesterRoles() []tempest.Snowflake {
	return t.requesterRoles
}

func (t *CommandContext) DeferResponse() error {
	return t.Defer(!t.isPrivateMessage)
}

func (t *CommandContext) EditResponse(message string) error {
	return t.EditReply(tempest.ResponseMessageData{
		Content: message,
	}, !t.isPrivateMessage)
}

func (t *CommandContext) SendRequesterPrivateMessage(msg string) (tempest.Message, error) {
	return t.tempest.SendPrivateMessage(t.requester, tempest.Message{
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
