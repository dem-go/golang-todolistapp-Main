package core_logger

import (
	"os"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger

	file *os.File
}
