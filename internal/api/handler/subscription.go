package handler

import (
	"github.com/danyaobertan/exchangemonitor/internal/subscriber"
	"net/http"
)

func (h *Handler) HandleSubscribe(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	if email == "" {
		http.Error(writer, "Email is required", http.StatusBadRequest)
		return
	}

	err := subscriber.SubscribeEmail(email)
	if err != nil {
		http.Error(writer, "Failed to subscribe email", http.StatusInternalServerError)
		return
	}

	_, _ = writer.Write([]byte("Email subscribed successfully"))
}
