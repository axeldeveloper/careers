package utils

import (
	"encoding/json"
	"net/http"
)

// @todo
// rset messages
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// @todo
// response json
func Respond(w http.ResponseWriter, statusCode int, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	//w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
