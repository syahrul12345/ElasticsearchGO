package scheduler

import (
	"cache/models"
	"cache/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Country struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
}

// Immediately load the scheduler when starting...
func init() {
	countries := getData()
	db := models.GetDB()
	fmt.Println(&db)
	schedulerTicker := time.NewTicker(4 * time.Second)
	go func() {
		for _, country := range countries {
			select {
			case <-schedulerTicker.C:
				fmt.Printf("Calling the api for country %s:%s \n", country.Name, country.Code)
				data := utils.CallDB(country.Code)
				db.Put([]byte(country.Code), *data, nil)
			}
		}
	}()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getData() []Country {
	data, err := ioutil.ReadFile("./scheduler/data.json")
	check(err)
	countries := &[]Country{}
	json.Unmarshal(data, countries)
	return *countries
}
