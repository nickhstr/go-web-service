package log

import (
	"github.com/nickhstr/go-web-service/app/utils/env"
	log "github.com/sirupsen/logrus"
)

func init() {
	formatter := jsonFormatter()
	log.SetFormatter(formatter)
}

func jsonFormatter() *log.JSONFormatter {
	formatter := &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}

	if !env.IsProd() {
		formatter.PrettyPrint = true
	}

	return formatter
}

// Leaving this here as another option
func textFormatter() *log.TextFormatter {
	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	if !env.IsProd() {
		formatter.ForceColors = true
		formatter.DisableLevelTruncation = true
	}

	return formatter
}
