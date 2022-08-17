package logger

type Logger interface {
	Info(string, map[string]interface{})
	Error(string, error)
}
