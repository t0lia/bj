package handler

import (
	"fmt"
	"math/rand"
	"net/http"
)

var defaultTasks = []string{
	"buy milk",
	"walk the dog",
	"reply to emails",
}

func NewRouter() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TODO:\n", getTask())
	})

	return mux
}

func getTask() string {
	return defaultTasks[rand.Intn(len(defaultTasks))]
}
