package main

import (
	"fmt"
	"log/slog"
	"os"

	"example.com/basics/errors/errors"
)

func fn1() error {
	err := fmt.Errorf("other error")
	return errors.Wrap(err, "fn1 error")
}

func fn2() error {
	err := fn1()
	if err != nil {
		return errors.Wrap(err, "fn2 error")
	}
	return nil
}

func main() {
	err := fn2()
	// handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(handler)

	if err != nil {
		// Flatten the error to get a detailed and readable stack trace
		logger.Error("error", slog.Any("err", errors.GetStackTrace(err)))
	}
}
