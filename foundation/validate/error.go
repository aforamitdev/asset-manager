package validate

type AppError struct {
	Message string `json:"message"`
	Err     error  `json:"error"`
}

func NewAppError(err error, message string) error {
	return &AppError{Message: message, Err: err}

}

func (err *AppError) Error() string {
	return err.Err.Error()
}
