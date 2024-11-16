package logs

import (
	"log/slog"
	"os"
)

func logDefault() (logger *slog.Logger) {
	// green := "[0;92m"
	// reset := "\033[0m"
	// customHandler := slog.NewTextHandler(slog.NewTex{
	// 	Level: slog.LevelInfo,
	// 	Format: func(record slog.Record) string {
	// 		// Se a mensagem de log contém "SUCCESS", adicione cor verde
	// 		if record.Message == "SUCCESS -> We have a winner!" {
	// 			return fmt.Sprintf("%s%s%s", green, record.Message, reset)
	// 		}
	// 		// Retorne a mensagem original para outros logs
	// 		return record.Message
	// 	},
	// })
	// logJson := slog.New(customHandler)
	// logger = logJson
	logJson := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logJson)
	logger = logJson
	return
}

var Log = logDefault()
