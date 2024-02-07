package api_gateway

import (
	"errors"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/ogen-go/ogen/validate"
)

func translateError(err error) error {
	var ogenErr *validate.UnexpectedStatusCodeError
	if errors.As(err, &ogenErr) {
		switch ogenErr.StatusCode {
		case 400:
			return gateway.ErrInvalidInputError
		case 401:
			return gateway.ErrUnauthorizedError
		case 403:
			return gateway.ErrUnauthorizedError
		case 404:
			return gateway.ErrResourceNotFoundError
		}
	}
	return gateway.ErrInternalServerError
}
