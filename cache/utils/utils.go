package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
	Country string
}

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

// CallDB : Makes a call to the webserver
func CallDB(country string) *[]byte {
	tempCountry := &Country{
		Country: country,
	}
	requestBody, err := json.Marshal(tempCountry)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	resp, err := http.Post("http://127.0.0.1:5555/api/v1/search", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &body
}
