package api_gateway

// func translateError(err error) error {
// 	var ogenErr *validate.UnexpectedStatusCodeError
// 	if errors.As(err, &ogenErr) {
// 		log.Info().Msgf("unexpected status code: %d", ogenErr.StatusCode)
// 		switch ogenErr.StatusCode {
// 		case 400:
// 			return gateway.ErrInvalidInputError
// 		case 401:
// 			return gateway.ErrUnauthorizedError
// 		case 403:
// 			return gateway.ErrUnauthorizedError
// 		case 404:
// 			return gateway.ErrResourceNotFoundError
// 		}
// 	}
// 	log.Error().Err(err).Msg("unexpected error")
// 	return gateway.ErrInternalServerError
// }
