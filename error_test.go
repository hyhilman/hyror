package pq_error

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
	"strings"
	"testing"
)

func printRightErrorMessage(t *testing.T, err error, msg string)  {
	if strings.Contains(fmt.Sprintf("%s", err), msg) {
		t.Logf("\033[1;34msuccess: contain error message %s\033[0m", msg)
	} else {
		t.Errorf("\033[1;31mfailed: should have %s but shown %s\033[0m", msg, err)
	}
}

func testStack(t *testing.T, err error, line int)  {
	testString := fmt.Sprintf("error_test.go:%d", line)

	if strings.Contains(fmt.Sprintf("%s", err), testString) {
		t.Logf("\033[1;34msuccess: have right error stacks\033[0m")
	} else {
		t.Errorf("\033[1;31mfailed: should have %s but shown %s\033[0m", testString, err)
	}
}

func testNoStack(t *testing.T, err error, line int)  {
	testString := fmt.Sprintf("error_test.go:%d", line)

	if !strings.Contains(fmt.Sprintf("%s", err), testString) {
		t.Logf("\033[1;34msuccess: have right error stacks\033[0m")
	} else {
		t.Errorf("\033[1;31mfailed: should have %s but shown %s\033[0m", testString, err)
	}
}

func TestSimpleError(t *testing.T) {
	msg := "with stack"
	simpleError := errors.New(msg)
	err := NewError(simpleError)

	if strings.Contains(fmt.Sprintf("%s", err), t.Name()) {
		t.Logf("\u001B[1;34msuccess: have %s error method name\033[0m", t.Name())
	} else {
		t.Errorf("\033[1;31mfailed: doesn't have %s error method\u001B[0m", t.Name())
	}

	printRightErrorMessage(t, err, msg)
	testStack(t, err, 42)
}

func TestNoErrorStackNotThrowingPanic(t *testing.T)  {
	noStackError := net.ErrWriteToConnected
	msg := "use of WriteTo with pre-connected connection"

	printRightErrorMessage(t, noStackError, msg)
	testNoStack(t, noStackError, 55)
}

func TestNoErrorStackHaveStackWhenCreatedProperly(t *testing.T)  {
	stackError := NewError(net.ErrWriteToConnected)
	msg := "use of WriteTo with pre-connected connection"

	printRightErrorMessage(t, stackError, msg)
	testStack(t, stackError, 63)
}

func TestGoRoutineError(t *testing.T)  {
	var err error
	errChan := make(chan error)

	go func(e chan error) {
		e <- NewError(net.ErrWriteToConnected)
	}(errChan)

	err = <- errChan
	msg := "use of WriteTo with pre-connected connection"

	printRightErrorMessage(t, err, msg)
	testStack(t, err, 75)
}

func TestStackPrintWithoutStack(t *testing.T)  {
	msg := "with stack"
	simpleError := errors.New(msg)
	err := NewError(simpleError)

	if fmt.Sprintf("%s", err.(PQError).ErrorWithoutStack()) == msg {
		t.Logf("\u001B[1;34msuccess: have error message %s without stack\033[0m", msg)
	} else {
		t.Errorf("\033[1;31mfailed: failed to hide stack, given: %s\u001B[0m", err.Error())
	}

	printRightErrorMessage(t, err, msg)
}