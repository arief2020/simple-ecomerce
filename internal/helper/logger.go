package helper

import (
	// "errors"
	// "fmt"
	// "runtime"

	// "github.com/rs/zerolog/log"
	// "errors"
	// "runtime"

	"github.com/sirupsen/logrus"
)

const (
	LoggerLevelTrace = "LoggerLevelTrace"
	LoggerLevelDebug = "LoggerLevelDebug"
	LoggerLevelInfo  = "LoggerLevelInfo"
	LoggerLevelWarn  = "LoggerLeveWarn"
	LoggerLevelError = "LoggerLevelError"
	LoggerLevelFatal = "LoggerLevelFatal"
	LoggerLevelPanic = "LoggerLevelPanic"
)

// func Logger2(level, message string, err error) {
// 	if err == nil && (level == "" || message == "") {
// 		log.Error().Stack().Err(errors.New("all params log is required")).Msg("")
// 	}

// 	pc, _, line, _ := runtime.Caller(1)
// 	path := runtime.FuncForPC(pc).Name()

//		switch level {
//		case LoggerLevelDebug:
//			log.Debug().Str("message", message).Msg("")
//		case LoggerLevelInfo:
//			log.Info().Str("message", message).Msg("")
//		case LoggerLevelWarn:
//			log.Warn().Str("message", message).Msg("")
//		case LoggerLevelError:
//			log.Error().Str("path", path).Str("line", fmt.Sprint(line)).Err(err).Send()
//		case LoggerLevelFatal:
//			log.Fatal().Str("path", path).Str("line", fmt.Sprint(line)).Err(err).Send()
//		case LoggerLevelPanic:
//			log.Panic().Str("path", path).Str("line", fmt.Sprint(line)).Err(err).Send()
//		default:
//			log.Error().Stack().Err(errors.New("logger level invalid")).Send()
//		}
//	}
func Logger(filepath, level, message string) {
	if filepath == "" || level == "" || message == "" {
		logrus.WithFields(
			logrus.Fields{
				"file": "internal/helper/logger.go",
			},
		).Error("All params is required")
	}

	logging := logrus.WithFields(
		logrus.Fields{
			"file": filepath,
		})

	switch level {
	case LoggerLevelDebug:
		logging.Debug(message)
	case LoggerLevelInfo:
		logging.Info(message)
	case LoggerLevelWarn:
		logging.Warn(message)
	case LoggerLevelError:
		logging.Error(message)
	case LoggerLevelFatal:
		logging.Fatal(message)
	case LoggerLevelPanic:
		logging.Panic(message)
	default:
		logging.Error("Level invalid")
	}

}
