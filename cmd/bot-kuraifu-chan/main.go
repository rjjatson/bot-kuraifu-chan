package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("bot-kuraifu-chan is running...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "fuwaaa~")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
