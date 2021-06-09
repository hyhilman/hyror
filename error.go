package pq_error

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErrorWithStackTrace interface {
	StackTrace() errors.StackTrace
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

func NewError(err error) error {
	return PQError{errors.Wrap(err, err.Error())}
}
