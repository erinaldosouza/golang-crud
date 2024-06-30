package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	/*logger, _ := zap.NewProduction()

	defer logger.Sync()
	logger.Info("Application failed",
		zap.String(" ", " "),
	)*/

	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLogLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "nivel",
			TimeKey:      "tempo",
			MessageKey:   "mensagem",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()

}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_OUTPUT")))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLogLevel() zapcore.Level {

	switch strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {

	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}

}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(message, tags...)
	log.Sync()
}
