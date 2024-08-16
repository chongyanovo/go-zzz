package logger

import "go.uber.org/zap"

type ZapLogger struct {
	l *zap.Logger
}

func NewZapLogger(l *zap.Logger) Logger {
	return &ZapLogger{l: l}
}

func (z *ZapLogger) Debug(msg string, fields ...LogField) {
	z.l.Debug(msg, z.toZapFields(fields)...)
}

func (z *ZapLogger) Info(msg string, fields ...LogField) {
	z.l.Info(msg, z.toZapFields(fields)...)
}

func (z *ZapLogger) Warn(msg string, fields ...LogField) {
	z.l.Warn(msg, z.toZapFields(fields)...)
}

func (z *ZapLogger) Error(msg string, fields ...LogField) {
	z.l.Error(msg, z.toZapFields(fields)...)
}

func (z *ZapLogger) Fatal(msg string, fields ...LogField) {
	z.l.Fatal(msg, z.toZapFields(fields)...)
}

func (z *ZapLogger) Panic(msg string, fields ...LogField) {
	z.l.Panic(msg, z.toZapFields(fields)...)
}

func (z *ZapLogger) toZapFields(fields []LogField) []zap.Field {
	res := make([]zap.Field, 0, len(fields))
	for _, field := range fields {
		res = append(res, zap.Any(field.Key, field.Value))
	}
	return res
}
