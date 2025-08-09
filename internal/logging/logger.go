package logging

type Field struct {
	Key   string
	Value interface{}
}

type Logger interface {
	Info(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

func NewField(key string, value interface{}) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}
