package logger

type Logger interface {
	Debug(msg string, fields ...LogField)
	Info(msg string, fields ...LogField)
	Warn(msg string, fields ...LogField)
	Error(msg string, fields ...LogField)
	Fatal(msg string, fields ...LogField)
	Panic(msg string, fields ...LogField)
}

type LogField struct {
	Key   string
	Value any
}

func NewLogField(key string, value any) LogField {
	return LogField{
		Key:   key,
		Value: value,
	}
}

func String(key string, value string) LogField {
	return NewLogField(key, value)
}

func Any(key string, value any) LogField {
	return NewLogField(key, value)
}

func Error(err error) LogField {
	return NewLogField("error", err)
}
