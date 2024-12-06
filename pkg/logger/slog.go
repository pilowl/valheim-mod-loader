package logger

import (
	"log/slog"
	"os"
	"path/filepath"
)

type Logger struct {
	*slog.Logger
}

func NewLogger() *Logger {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	return &Logger{slog.New((slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replace})))}
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		l.Logger.With(args),
	}
}

func (l *Logger) WithError(err error) *Logger {
	return l.With("error", err)
}
