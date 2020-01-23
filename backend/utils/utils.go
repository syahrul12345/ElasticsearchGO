package utils

import (
	"encoding/json"
	"net/http"
)

//Message outputs the message to be sent to the client
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
	}
}

//Respond automatically writes to the HTTP response, and will be displayed for the frontend
func Respond(writer http.ResponseWriter, message map[string]interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(message)
}
