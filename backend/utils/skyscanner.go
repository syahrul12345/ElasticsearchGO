package utils

import (
	"encoding/json"
)

//Country object
type Country struct {
	Country string `json:"country"`
}

//Routes object
type Routes struct {
	Routes []route
}
type route struct {
	OriginID      int     `json:"OriginId"`
	DestinationID int     `json:"DestinationId"`
	QuoteID       []int   `json:"QuoteIds"`
	Price         float64 `json:"Price"`
	QuoteDateTime string  `json:"QuoteDateTime"`
}

const (
	// URLENDPOINT is the endpoint to the string
	URLENDPOINT string = "http://partners.api.skyscanner.net/apiservices/browseroutes/v1.0/"
	URLQUERY    string = "/sgd/en-US/us/anywhere/anytime/anytime?apikey=prtl6749387986743898559646983194"
)

// GetRoutes returns all the routes for the given country
func (country *Country) GetRoutes() (map[string]interface{}, bool) {
	resp := make(map[string]interface{})
	response := Request(URLENDPOINT + country.Country + URLQUERY)
	routes := &Routes{}
	json.Unmarshal(*response, routes)
	if len(routes.Routes) == 0 {
		resp["error"] = "Error getting data from api"
		return resp, false
	}
	resp["routeList"] = routes
	return resp, true
}
