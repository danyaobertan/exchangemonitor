package utils

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"go.uber.org/zap"
)

func DecodeJSONRequest(r *http.Request, dest any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(dest)
}

func WriteJSONResponse(w http.ResponseWriter, l logger.Logger, statusCode int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		l.Error("Error writing response", zap.Error(err))
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

// ValidEmail checks if the given string is a valid email address.
func ValidEmail(email string) bool {
	// Regex for validating an Email
	regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}
