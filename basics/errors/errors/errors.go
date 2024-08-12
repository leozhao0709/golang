package errors

import (
	"errors"
	"fmt"
)

func Wrap(err error, msg string) stackTraceError {
	return newStackTraceError(fmt.Errorf("%w; %s", err, msg))
}

func GetStackTrace(err error) string {
	var stErr stackTraceError
	if ok := errors.As(err, &stErr); ok {
		return stErr.errorWithStack()
	}
	return err.Error()
}
