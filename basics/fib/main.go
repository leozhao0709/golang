package main

import (
	"log/slog"
	"os"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	slog := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	slog.Info("pid", "pid", os.Getpid())
	res := fib(48)
	slog.Info("res: ", "res", res)
}
