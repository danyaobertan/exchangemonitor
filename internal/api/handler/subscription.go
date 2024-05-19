package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/danyaobertan/exchangemonitor/models"
	"github.com/danyaobertan/exchangemonitor/pkg/utils"
	"go.uber.org/zap"
)

type SubscribeUserRequestBody struct {
	Email string `json:"email"`
}

func (s SubscribeUserRequestBody) Validate() error {
	if s.Email == "" && !utils.ValidEmail(s.Email) {
		return errors.New("email is required")
	}

	return nil
}

type SubscribeResponseBody struct {
	Message string `json:"message"`
}

// HandleSubscribe handles subscription requests
func (h *Handler) HandleSubscribe(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var requestBody SubscribeUserRequestBody
	if err := utils.DecodeJSONRequest(request, &requestBody); err != nil {
		h.log.Error("Failed to decode request body", zap.Error(err))
		http.Error(writer, "Invalid request body", http.StatusBadRequest)

		return
	}

	if err := requestBody.Validate(); err != nil {
		h.log.Error("Invalid request body", zap.Error(err))
		http.Error(writer, "Invalid request body", http.StatusBadRequest)

		return
	}

	subscriber, err := h.dbClient.GetSubscription(ctx, models.Subscriber{Email: requestBody.Email})
	if err == nil && subscriber.Email == requestBody.Email {
		h.log.Info("Email already subscribed", zap.String("email", subscriber.Email))
		http.Error(writer, fmt.Sprintf("Email %s already subscribed", subscriber.Email), http.StatusConflict)

		return
	}

	if err = h.dbClient.AddSubscription(ctx, models.Subscriber{Email: requestBody.Email}); err != nil {
		h.log.Error("Failed to subscribe email", zap.Error(err))
		http.Error(writer, "Failed to subscribe email", http.StatusInternalServerError)

		return
	}

	resp, err := json.Marshal(SubscribeResponseBody{Message: fmt.Sprintf("Email %s subscribed", requestBody.Email)})
	if err != nil {
		h.log.Error("Unable to marshal response: %v", zap.Error(err))
		http.Error(writer, "Unable to marshal response", http.StatusInternalServerError)

		return
	}

	utils.WriteJSONResponse(writer, h.log, http.StatusOK, resp)
}
