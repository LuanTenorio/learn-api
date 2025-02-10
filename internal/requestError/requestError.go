package requestError

const (
	IncompatibleBody = "The body is incompatible with the desired"
)

type RequestErrorBody struct {
	Message string
}

func New(message string) *RequestErrorBody {
	return &RequestErrorBody{Message: message}
}
