package logger

type Logger interface {
	Log(data string)
	Error(data string)
	Debug(data string)
}
