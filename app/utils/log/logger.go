package log

import (
	"github.com/nickhstr/go-web-service/app/utils/env"
	log "github.com/sirupsen/logrus"
)

func init() {
	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	if !env.IsProd() {
		formatter.ForceColors = true
		formatter.DisableLevelTruncation = true
	}

	log.SetFormatter(formatter)
}
