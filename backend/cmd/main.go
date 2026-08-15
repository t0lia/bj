package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/t0lia/bj/internal/handler"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	router := handler.NewRouter()

	slog.Info("server listening", "addr", ":8080")
	http.ListenAndServe(":8080", router)
}
