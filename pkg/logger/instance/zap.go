package instance

import (
	"github.com/IamStubborN/test/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger(level, encodingType string, outputs []string) (logger.Logger, error) {
	var lvl zapcore.Level
	if err := lvl.Set(level); err != nil {
		return nil, err
	}

	l, err := zap.Config{
		Encoding:    encodingType,
		Level:       zap.NewAtomicLevelAt(lvl),
		OutputPaths: outputs,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}.Build()

	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(l)

	return &zapLogger{logger: l.Sugar()}, nil
}

func (zl *zapLogger) Info(data ...interface{}) {
	zl.logger.Info(data...)
}

func (zl *zapLogger) Warn(data ...interface{}) {
	zl.logger.Warn(data...)
}

func (zl *zapLogger) Fatal(data ...interface{}) {
	zl.logger.Fatal(data...)
}

func (zl *zapLogger) WithFields(data []interface{}) logger.Logger {
	return &zapLogger{logger: zl.logger.With(data...)}
}
