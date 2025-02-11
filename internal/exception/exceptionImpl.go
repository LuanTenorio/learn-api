package exception

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	IncompatibleBody = "the body is incompatible with the desired"
)

type ExceptionImpl struct {
	Message       string   `json:"message"`
	Code          int      `json:"code"`
	Trace         []string `json:"-"`
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

func (rrb *ExceptionImpl) AddTraceLog(info string) {
	rrb.Trace = append(rrb.Trace, rrb.getInfoFromLastCallStack()+info)
}

func (rrb *ExceptionImpl) Error() string {
	return rrb.Message
}

func (rrb *ExceptionImpl) HttpException(c echo.Context) error {
	return c.JSON(rrb.Code, rrb)
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
