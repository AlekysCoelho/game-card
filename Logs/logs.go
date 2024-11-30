package logs

import (
	"log/slog"
	"os"
)

func logDefault() (logger *slog.Logger) {
	logJson := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logJson)
	logger = logJson
	return
}

var Log = logDefault()
