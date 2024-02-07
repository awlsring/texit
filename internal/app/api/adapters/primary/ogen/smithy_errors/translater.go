package smithy_errors

import (
	"errors"

	"github.com/awlsring/texit/internal/app/api/ports/service"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
)

func translateError(err error) *SmithyError {
	var (
		ctError *validate.InvalidContentTypeError
		ogenErr ogenerrors.Error
	)
	switch {
	case errors.Is(err, ht.ErrNotImplemented):
		return NewUnknownOperationError(err.Error())
	case errors.Is(err, service.ErrUnknownNode) || errors.Is(err, service.ErrUnknownPlatform) || errors.Is(err, service.ErrUnknownTailnetId) || errors.Is(err, service.ErrUnknownProvider) || errors.Is(err, service.ErrUnknownTailnet) || errors.Is(err, service.ErrExecutionNotFound):
		return NewResourceNotFoundError(err.Error())
	case errors.As(err, &ctError):
		return NewInternalFailureError(err.Error())
	case errors.As(err, &ogenErr):
		code := ogenErr.Code()
		switch code {
		case 400:
			return NewInvalidInputError(err.Error())
		}
	}

	return NewInternalFailureError(err.Error())
}
