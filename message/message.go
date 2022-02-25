package message

import (
	"fmt"
	"net/http"
)

type Respones struct {
	Message string `json:"message"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("send message")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))

	w.Header().Set("Content-Type", "application/json")
}
