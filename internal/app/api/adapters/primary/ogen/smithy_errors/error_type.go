package smithy_errors

type ErrorType int

const (
	ErrorTypeInternalServerError ErrorType = iota
	ErrorTypeUnauthorizedError
	ErrorTypeInvalidInputError
	ErrorTypeResourceNotFoundError
	ErrorTypeInternalFailureError
	ErrorTypeSerializationError
	ErrorTypeUnknownOperationError
	ErrorTypeValidationError
)

func (e ErrorType) String() string {
	switch e {
	case ErrorTypeInternalServerError:
		return "InternalServerError"
	case ErrorTypeUnauthorizedError:
		return "UnauthorizedError"
	case ErrorTypeInvalidInputError:
		return "InvalidInputError"
	case ErrorTypeResourceNotFoundError:
		return "ResourceNotFoundError"
	case ErrorTypeSerializationError:
		return "SerializationError"
	case ErrorTypeUnknownOperationError:
		return "UnknownOperationError"
	case ErrorTypeValidationError:
		return "ValidationError"
	default:
		return "InternalFailureError"
	}
}
