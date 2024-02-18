package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/time", TimeHandler)

	if err := http.ListenAndServe(":8795", nil); err != nil {
		panic(err)
	}
}
