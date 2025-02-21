package exception

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const (
	IncompatibleBody = "the body is incompatible with the desired"
)

type ExceptionImpl struct {
	Message       string `json:"message"`
	Code          int    `json:"code"`
	trace         []string
	indexFromCall int
}

func New(message string, code int, logs ...string) *ExceptionImpl {
	exc := &ExceptionImpl{Message: message, Code: code, indexFromCall: 3}

	if len(logs) == 0 {
		exc.AddTraceLog(message)
	} else {
		for _, log := range logs {
			exc.AddTraceLog(log)
		}
	}

	exc.indexFromCall--

	return exc
}

func NewCanceledRequest(message string) *ExceptionImpl {
	return New("Request canceled", http.StatusBadRequest, message)
}

func (rrb *ExceptionImpl) AddTraceLog(info string) Exception {
	rrb.trace = append(rrb.trace, rrb.getInfoFromLastCallStack()+info)
	return rrb
}

func (rrb *ExceptionImpl) Error() string {
	return rrb.Message
}

func (rrb *ExceptionImpl) HttpException(c echo.Context) error {
	return c.JSON(rrb.Code, rrb)
}

func (rrb *ExceptionImpl) GetTrace() []string {
	return rrb.trace
}

func (rrb *ExceptionImpl) getInfoFromLastCallStack() string {
	_, file, line, ok := runtime.Caller(rrb.indexFromCall)

	if !ok {
		fmt.Println("It was not possible to obtain information from the Caller")
		return ""
	}

	files := strings.Split(file, "/")
	errorFileName := files[len(files)-1]

	return fmt.Sprintf("%s - %d: ", errorFileName, line)
}

func CheckExceptionForTest(t *testing.T, err error, expectedException ExceptionImpl) {
	assert.Implements(t, (*Exception)(nil), err)
	if exeption, ok := err.(*ExceptionImpl); assert.True(t, ok) {
		assert.Equal(t, exeption.Code, expectedException.Code)
		assert.Equal(t, exeption.Message, expectedException.Message)
	}
}

func CheckDbException(err error) Exception {
	if err == nil {
		return nil
	} else if errors.Is(err, context.Canceled) {
		return NewCanceledRequest(err.Error())
	}

	return New("Db internal error", http.StatusInternalServerError, err.Error())
}
