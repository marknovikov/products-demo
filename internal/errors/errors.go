package errors

type ErrInvalidInput struct {
	Base error
}

func NewErrInvalidInput(base error) error {
	return ErrInvalidInput{Base: base}
}

func (err ErrInvalidInput) Error() string {
	return "InvalidInput: " + err.Base.Error()
}

func (err ErrInvalidInput) Unwrap() error {
	return err.Base
}

type ErrInternal struct {
	Base error
}

func NewErrInternal(base error) error {
	return ErrInternal{Base: base}
}

func (err ErrInternal) Error() string {
	return "Internal: " + err.Base.Error()
}

func (err ErrInternal) Unwrap() error {
	return err.Base
}
