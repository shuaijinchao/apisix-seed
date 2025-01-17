package log

var (
	DefLogger Interface = emptyLog{}
)

type Type int8

type emptyLog struct {
}

type Interface interface {
	Debug(msg string, fields ...interface{})
	Debugf(msg string, args ...interface{})
	Info(msg string, fields ...interface{})
	Infof(msg string, args ...interface{})
	Warn(msg string, fields ...interface{})
	Warnf(msg string, args ...interface{})
	Error(msg string, fields ...interface{})
	Errorf(msg string, args ...interface{})
	Fatal(msg string, fields ...interface{})
	Fatalf(msg string, args ...interface{})
}

func (e emptyLog) Debug(msg string, fields ...interface{}) {
	getZapFields(logger, fields).Debug(msg)
}

func (e emptyLog) Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

func (e emptyLog) Info(msg string, fields ...interface{}) {
	getZapFields(logger, fields).Info(msg)
}

func (e emptyLog) Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

func (e emptyLog) Warn(msg string, fields ...interface{}) {
	getZapFields(logger, fields).Warn(msg)
}

func (e emptyLog) Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

func (e emptyLog) Error(msg string, fields ...interface{}) {
	getZapFields(logger, fields).Error(msg)
}

func (e emptyLog) Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

func (e emptyLog) Fatal(msg string, fields ...interface{}) {
	getZapFields(logger, fields).Fatal(msg)
}

func (e emptyLog) Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}

func Debug(msg string, fields ...interface{}) {
	DefLogger.Debug(msg, fields...)
}
func Debugf(msg string, args ...interface{}) {
	DefLogger.Debugf(msg, args...)
}
func Info(msg string, fields ...interface{}) {
	DefLogger.Info(msg, fields...)
}
func Infof(msg string, args ...interface{}) {
	DefLogger.Infof(msg, args...)
}
func Warn(msg string, fields ...interface{}) {
	DefLogger.Warn(msg, fields...)
}
func Warnf(msg string, args ...interface{}) {
	DefLogger.Warnf(msg, args...)
}
func Error(msg string, fields ...interface{}) {
	DefLogger.Error(msg, fields...)
}
func Errorf(msg string, args ...interface{}) {
	DefLogger.Errorf(msg, args...)
}
func Fatal(msg string, fields ...interface{}) {
	DefLogger.Fatal(msg, fields...)
}
func Fatalf(msg string, args ...interface{}) {
	DefLogger.Fatalf(msg, args...)
}
