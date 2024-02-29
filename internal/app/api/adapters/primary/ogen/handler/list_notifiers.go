package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) ListNotifiers(ctx context.Context) (*texit.ListNotifiersResponseContent, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved list notifier request")

	log.Debug().Msg("Listing notifiers")
	notifiers := h.notifierSvc.ListNotifiers(ctx)
	log.Debug().Msgf("Found %d notifiers", len(notifiers))
	if len(notifiers) == 0 {
		log.Warn().Msg("No notifiers found")
		return &texit.ListNotifiersResponseContent{}, nil
	}

	log.Debug().Msg("Converting notifier to summaries")
	summaries := make([]texit.NotifierSummary, len(notifiers))
	for i, notifier := range notifiers {
		summaries[i] = conversion.NotifierToSummary(notifier)
	}

	log.Debug().Msg("Successfully listed notifiers")
	return &texit.ListNotifiersResponseContent{
		Summaries: summaries,
	}, nil
}
