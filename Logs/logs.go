package logs

import (
	"log/slog"
	"os"
)

func LogDefault() (logger *slog.Logger) {
	logJson := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logJson)
	logger = logJson
	return
}
