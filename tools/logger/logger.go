package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log        *zap.Logger
	LOG_OUTPUT = os.Getenv("LOG_OUTPUT")
	LOG_LEVEL  = os.Getenv("LOGLEVEL")
)

// Função que instãncia o Zap para configurar o log da aplicação
func init() {
	logConfiguration := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "tempo",
			NameKey:      "name",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfiguration.Build()
}

// Função para mostrar log de informações da aplicação
func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

// Função para mostrar log de erros da aplicação
func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(message, tags...)
	log.Sync()
}

// Função para mostrar log de alertas da aplicação
func Warn(message string, tags ...zap.Field) {
	log.Warn(message, tags...)
	log.Sync()
}

// Função que retorna o tipo de saída do log
func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(LOG_OUTPUT))
	if output == "" {
		return "stdout"
	}

	return output
}

// Função que retorna o nível de log
func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(LOG_LEVEL)) {
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
