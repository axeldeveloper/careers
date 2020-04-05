package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	//err := json.NewEncoder(w).Encode(data)
	json.NewEncoder(w).Encode(err)

}
