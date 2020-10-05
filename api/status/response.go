package status

import (
	"encoding/json"
	"net/http"
)

var validationError string = "Validation Error"
var serverError string = "Internal Server Error"
var authorizationError string = "Not Authorized"

// RespondStatus -
func RespondStatus(w http.ResponseWriter, code int, err error) {
	errorMap := map[string]string{}
	if err != nil {
		errorMap["Error"] = err.Error()
	}
	_, ok := errorMap["Error"]
	w.WriteHeader(code)
	switch code {
	case 400:
		if !ok {
			errorMap["Error"] = validationError
		}
		json.NewEncoder(w).Encode(&errorMap)
	case 401:
		if !ok {
			errorMap["Error"] = authorizationError
		}
		json.NewEncoder(w).Encode(&errorMap)
	case 500:
		if !ok {
			errorMap["Error"] = serverError
		}
		json.NewEncoder(w).Encode(&errorMap)
	default:
		errorMap["Error"] = serverError
		json.NewEncoder(w).Encode(&errorMap)
	}
}
