package controller

import (
	"cache/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

type route struct {
	OriginID      int     `json:"OriginId"`
	DestinationID int     `json:"DestinationId"`
	QuoteID       []int   `json:"QuoteIds"`
	Price         float64 `json:"Price"`
	QuoteDateTime string  `json:"QuoteDateTime"`
}

type RouteList struct {
	Routes []route
}

type ResponseBody struct {
	RouteList `json:"routeList"`
}
type Country struct {
	Country string
}

// GetSearch Attemps to get the expercted result from the database when provided with a search term.
var GetSearch = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	/*
		Get the stored value for the search term
		payload must be in the format of
		{
			country:"country"
		}
	*/

	tempCountry := &Country{}
	err := json.NewDecoder(r.Body).Decode(tempCountry)
	fmt.Printf("Received a request to get data from the cahce for country %s\n", tempCountry.Country)
	if err != nil {
		fmt.Println("Failed to parse incoming payload...")
		// Handle a generic error
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}

	// Open the db
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	data, err := db.Get([]byte(tempCountry.Country), nil)
	respBody := &ResponseBody{}
	if err != nil {
		fmt.Printf("Data for %s not in cahce. Calling from skyscanner api\n", tempCountry.Country)
		// It doesnt exist in the database. Make a call to the main server
		data := utils.CallDB(tempCountry.Country)
		// Save the byte array in our leveldb
		db.Put([]byte(tempCountry.Country), *data, nil)
		json.Unmarshal(*data, respBody)
	} else {
		fmt.Printf("Data for %s in cache.\n", tempCountry.Country)
		json.Unmarshal(data, respBody)
	}
	if len(respBody.RouteList.Routes) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}
	resp["routeList"] = respBody.RouteList
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, resp)
	return

}
