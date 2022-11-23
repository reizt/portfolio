package core

var (
	ErrInvalidInput = NewCoreError("invalid input")
	ErrUnexpected   = NewCoreError("unexpected")
)

type coreError struct {
	sentence string
}

func (e *coreError) Error() string {
	return e.sentence
}

func NewCoreError(text string) error {
	return &coreError{text}
}
