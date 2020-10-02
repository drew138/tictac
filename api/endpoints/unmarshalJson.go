package endpoints

import (
	"encoding/json"
	"net/http"
)

// UnmarshalJSON - apply json body of a request context to an especified model
func UnmarshalJSON(r *http.Request, model interface{}) error {
	err := json.Unmarshal([]byte(r.Body, model))
	return err
}
