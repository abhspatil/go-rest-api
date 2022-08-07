package exception

type error interface {
	Error() string
}

type GenericError struct {
	Msg string `json:"msg"` // description of error
}

func (e *GenericError) SetError(errMsg string) {
	e.Msg = errMsg
}

func (e *GenericError) Error() string { return e.Msg }
