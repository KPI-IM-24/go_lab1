package main

import (
	"encoding/json"
	"lab1/model"
	"net/http"
	"time"
)

func TimeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	response := model.TimeResponse{
		Time: time.Now().Format(time.RFC3339),
	}

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
