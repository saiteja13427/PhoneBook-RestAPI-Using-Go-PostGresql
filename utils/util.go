package utils

import (
	"encoding/json"
	"net/http"
)
//Message and Respond functions are handy functions for building and sending messages
func Message(status bool, message string) map[string]interface{}{
	return map[string]interface{} {"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}