package handler

import (
	"github.com/danyaobertan/exchangemonitor/pkg/currency"
	"net/http"
	"strconv"
)

func (h *Handler) HandleGetRate(writer http.ResponseWriter, request *http.Request) {
	rate, err := currency.FetchCurrentRateNBU()
	if err != nil {
		http.Error(writer, "Unable to fetch rate", http.StatusBadRequest)
		return
	}
	message := "Rate: " + strconv.FormatFloat(rate, 'f', -1, 64)
	_, _ = writer.Write([]byte(message))
}
