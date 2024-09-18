package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello PSS.sk! %s\n", hostname)
	})

	fmt.Println("Listen on 0.0.0.0:8000, see: http://127.0.0.1:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
