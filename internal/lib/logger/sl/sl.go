package sl

import (
	"log/slog"
)

func Err(err string) *slog.Attr {
	return &slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err),
	}
}
