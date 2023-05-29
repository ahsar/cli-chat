// for logger with traceID
// usage:
// logger.Init("info")
// logger.InfoC(ctx, "msg")
package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(logLevel string) {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
		NoColor:    false,
	})

	switch strings.ToLower(logLevel) {
	case zerolog.DebugLevel.String(): // debug
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case zerolog.InfoLevel.String(): // info
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case zerolog.WarnLevel.String(): // warn
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case zerolog.ErrorLevel.String(): // error
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case zerolog.FatalLevel.String(): // fatal
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case zerolog.PanicLevel.String(): // panic
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case zerolog.NoLevel.String(): // ""
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	case zerolog.Disabled.String(): // disabled
		zerolog.SetGlobalLevel(zerolog.Disabled)
	case zerolog.TraceLevel.String(): // trace
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func Debug(f string, args ...any) {
	log.Debug().Msgf(f, args...)
}

func Info(f string, args ...any) {
	log.Info().Msgf(f, args...)
}

func Warn(f string, args ...any) {
	log.Warn().Msgf(f, args...)
}

func Error(f string, args ...any) {
	log.Error().Msgf(f, args...)
}

func Fatal(f string, args ...any) {
	log.Fatal().Msgf(f, args...)
}

func Panic(f string, args ...any) {
	log.Panic().Msgf(f, args...)
}

//type SeverityHook struct{}

//func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
//if _, file, line, ok := runtime.Caller(5); ok {
//e.Str("line", fmt.Sprintf("%s:%d", file, line))
//}
//}
