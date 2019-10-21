package logger

type Logger interface {
	Info(...interface{})
	Warn(...interface{})
	Fatal(...interface{})
	WithFields([]interface{}) Logger
}
