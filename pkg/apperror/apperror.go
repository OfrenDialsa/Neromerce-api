package apperror

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

func New(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}
