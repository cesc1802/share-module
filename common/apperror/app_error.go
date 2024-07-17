package apperror

type AppError interface {
}

type appError struct {
	root error
}

func (e *appError) Error() string {

}
