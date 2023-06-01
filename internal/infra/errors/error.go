package errors

type (
	Err struct {
		msg string
	}
)

func NewError(msg string) Err {
	return Err{msg}
}

func (e Err) Error() string {
	return e.msg
}
