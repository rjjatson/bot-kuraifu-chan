package main

import (
	"net/http"
	"os"
)

func main() {
	healthPort := os.Getenv("HEALTH_PORT")
	if healthPort == "" {
		healthPort = "8888"
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	err := http.ListenAndServe(":"+healthPort, nil)
	if err != nil {
		panic(err)
	}
}
