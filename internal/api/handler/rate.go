package handler

import (
	"encoding/json"
	"github.com/danyaobertan/exchangemonitor/pkg/currency"
	"github.com/danyaobertan/exchangemonitor/pkg/utils"
	"go.uber.org/zap"
	"net/http"
)

type RateResponseBody struct {
	Rate float64 `json:"rate"`
}

// HandleGetRate handles rate requests
func (h *Handler) HandleGetRate(writer http.ResponseWriter, _ *http.Request) {
	rate, err := currency.FetchCurrentRateNBU()
	if err != nil {
		// На мою думку, 500 коректніше повертати, ніж 400 для цього випадку, попрои умову в завданні,
		//адже це помилка на сервері та користувач ніяк не може вплинути на її виправлення.
		h.log.Error("Unable to fetch rate: %v", zap.Error(err))
		http.Error(writer, "Unable to fetch rate", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(RateResponseBody{Rate: rate})
	if err != nil {
		h.log.Error("Unable to marshal response: %v", zap.Error(err))
		http.Error(writer, "Unable to marshal response", http.StatusInternalServerError)
		return
	}

	utils.WriteJSONResponse(writer, h.log, http.StatusOK, resp)
}
