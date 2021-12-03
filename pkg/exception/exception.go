package exception

import (
	"fmt"
)

type CustomError struct {
	errorString string
}

var (
	FAILURE_ERR       = CustomErrWrapper(FAILURE_MSG)
	INVALID_PARAM_ERR = CustomErrWrapper(INVALID_PARAM_MSG)
	SERVER_ERROR_ERR  = CustomErrWrapper(SERVER_ERROR_MSG)
)

func CustomErrWrapper(err string, args ...interface{}) *CustomError {
	if err == "" {
		return nil
	}
	return &CustomError{fmt.Sprintf(err, args...)}
}

func (e *CustomError) Error() string {
	if e == nil {
		return ""
	}
	return e.errorString
}
