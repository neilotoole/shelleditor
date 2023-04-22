package logging

import (
	"context"

	"golang.org/x/exp/slog"
)

// DefaultLogger is the default logger to use. Set it to a non-nil
// value if you want log output.
var DefaultLogger *slog.Logger

// Logger returns a non-nil *slog.Logger. If DefaultLogger is nil, Logger
// returns a *slog.Logger that discards its output.
func Logger() *slog.Logger {
	if DefaultLogger != nil {
		return DefaultLogger
	}

	return discard()
}

// discard returns a new *slog.Logger that discards output.
func discard() *slog.Logger {
	h := discardHandler{}
	return slog.New(h)
}

var _ slog.Handler = (*discardHandler)(nil)

type discardHandler struct{}

// Enabled implements slog.Handler.
func (d discardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

// Handle implements slog.Handler.
func (d discardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

// WithAttrs implements slog.Handler.
func (d discardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return d
}

// WithGroup implements slog.Handler.
func (d discardHandler) WithGroup(_ string) slog.Handler {
	return d
}
