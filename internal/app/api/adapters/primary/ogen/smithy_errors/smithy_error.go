package smithy_errors

type SmithyError struct {
	HttpCode  int
	ErrorType ErrorType
	Message   string
}

func (e *SmithyError) Error() string {
	return e.Message
}

func (e *SmithyError) Code() int {
	return e.HttpCode
}

func (e *SmithyError) Type() ErrorType {
	return e.ErrorType
}

func (e *SmithyError) AsJsonMessage() []byte {
	return []byte(`{"message": "` + e.Message + `"}`)
}

func NewUnauthorizedError(msg string) *SmithyError {
	return &SmithyError{
		HttpCode:  401,
		ErrorType: ErrorTypeUnauthorizedError,
		Message:   msg,
	}
}

func NewInvalidInputError(msg string) *SmithyError {
	return &SmithyError{
		HttpCode:  400,
		ErrorType: ErrorTypeInvalidInputError,
		Message:   msg,
	}
}

func NewResourceNotFoundError(msg string) *SmithyError {
	return &SmithyError{
		HttpCode:  404,
		ErrorType: ErrorTypeResourceNotFoundError,
		Message:   msg,
	}
}

func NewInternalFailureError(msg string) *SmithyError {
	return &SmithyError{
		HttpCode:  500,
		ErrorType: ErrorTypeInternalFailureError,
		Message:   msg,
	}
}

func NewSerializationError(msg string) *SmithyError {
	return &SmithyError{
		HttpCode:  400,
		ErrorType: ErrorTypeSerializationError,
		Message:   msg,
	}
}

func NewUnknownOperationError(msg string) *SmithyError {
	return &SmithyError{
		HttpCode:  400,
		ErrorType: ErrorTypeUnknownOperationError,
		Message:   msg,
	}
}
