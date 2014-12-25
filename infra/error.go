package infra

type Error struct {
	Status  int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(status int, msg string) *Error {
	return &Error{
		Status:  status,
		Message: msg,
	}
}
