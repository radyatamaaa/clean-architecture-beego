package logger

// NoopLogger is no-op version of logger.Logger.
type NoopLogger struct{}

// Debug implements Logger.
func (NoopLogger) Debug(string, ...interface{}) {}

// Info implements Logger.
func (NoopLogger) Info(string, ...interface{}) {}

// Warn implements Logger.
func (NoopLogger) Warn(string, ...interface{}) {}

// Error implements Logger.
func (NoopLogger) Error(string, ...interface{}) {}
