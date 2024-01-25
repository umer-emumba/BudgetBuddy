package config

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func logFilename() string {
	t := time.Now()
	dir := "logs"
	filename := fmt.Sprintf("%s/logfile_%s.log", dir, t.Format("2006-01-02"))

	// Create the logs directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to create logs directory:", err)
		}
	}

	return filename
}

func InitLogger() {

	// Configure Zap logger with file output and daily rotation
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.OutputPaths = []string{"stdout", logFilename()}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	// Create a logger
	var err error
	Logger, err = config.Build()
	if err != nil {
		fmt.Println("Failed to initialize Zap logger:", err)
	}
	defer Logger.Sync() // flushes buffer, if any
}
