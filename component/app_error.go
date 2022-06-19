package component

type AppError struct {
	Type    string
	Message string
	Reason  string
}

func (err *AppError) Error() string {
	return err.Message
}

func NewAppError(errorMessage string, errorType string, reason string) *AppError {
	return &AppError{Message: errorMessage, Type: errorType, Reason: reason}
}
