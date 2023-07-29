package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"

	"github.com/gorobot-nz/tengu-bot/internal"
)

var logger *zap.Logger

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	logLevel := os.Getenv("LOG_LEVEL")
	atoi, err := strconv.Atoi(logLevel)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert log level. Error: %s", err.Error()))
	}

	logger = internal.NewLogger(zapcore.Level(atoi))
}

func main() {

}
