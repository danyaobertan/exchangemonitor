package utils

import (
	"encoding/json"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"regexp"
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

//func ValidEmail(email string) bool {
//	// This is a very basic email validation, it checks if the email has an @ symbol
//	// and at least one dot in the domain part
//	return len(email) > 3 && strings.Contains(email, "@") && strings.Contains(email, ".")
//}

// ValidEmail checks if the given string is a valid email address.
func ValidEmail(email string) bool {
	// Regex for validating an Email
	regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}
