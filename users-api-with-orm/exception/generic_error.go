package exception

type error interface {
	Error() string
}

type GenericError struct {
	msg string // description of error
}

func (e *GenericError) Error() string { return e.msg }
