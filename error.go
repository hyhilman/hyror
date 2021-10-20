package hyror

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErrorWithStackTrace interface {
	StackTrace() errors.StackTrace
	Error() string
}

type PQError struct {
	e error
}

func (p PQError) Error() string {
	err, ok := p.e.(ErrorWithStackTrace)
	if !ok {
		return fmt.Sprintf("%s", p.e)
	}
	return fmt.Sprintf("%s%+v", errors.Cause(p.e), err.StackTrace())
}

func (p PQError) ErrorWithoutStack() string {
	return errors.Cause(p.e).Error()
}

func NewError(e interface{}) error {
	if msg, ok := e.(string); ok {
		return PQError{e: errors.New(msg)}
	}

	if err, ok := e.(PQError); ok {
		return err
	}

	if err, ok := e.(ErrorWithStackTrace); ok {
		return PQError{e: err}
	}

	return PQError{errors.New(fmt.Sprintf("%+v", e))}
}
