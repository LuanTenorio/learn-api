package logger

import "fmt"

type LoggerImpl struct {
	pkg string
	ctx string
}

var loggers = make(map[string]*LoggerImpl)

// context within the application
// pkg is the package (handler, usecase ...)
func New(ctx string, pkg string) Logger {
	key := createLoggerKey(ctx, pkg)

	logger, ok := loggers[key]

	if !ok {
		logger = &LoggerImpl{pkg: pkg, ctx: ctx}
		loggers[key] = logger
	}

	return logger
}

func createLoggerKey(ctx string, pkg string) string {
	return fmt.Sprintf("%s-%s", ctx, pkg)
}

func (l *LoggerImpl) Log(data string) {
	fmt.Printf("\033[32m[Log] [%s-%s]\033[0m %s\n", l.ctx, l.pkg, data)
}

func (l *LoggerImpl) Error(data string) {
	fmt.Printf("\033[31m[Error] [%s-%s]\033[0m %s\n", l.ctx, l.pkg, data)
}

func (l *LoggerImpl) Debug(data string) {
	fmt.Printf("\033[33m[Debug] [%s-%s]\033[0m %s\n", l.ctx, l.pkg, data)
}
