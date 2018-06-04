package dialogflow

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ujjkumsi/docker-go/util"
)

type dialogResponse struct {
	fulfillmentText string `json:"fulfillmentText"`
}

type params struct {
	City   string `json:"city"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

/*
DialogflowHandler Handles the communication with the dialog flow
And returns proper response
*/
func DialogflowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(util.FormatRequest(r))
	var err error
	var dfr *Request
	// var p params

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&dfr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Println(dfr.QueryResult.QueryText)
	// Filter on action, using a switch for example

	// Retrieve the params of the request
	// if err = dfr.GetParams(&p); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// Retrieve a specific context
	// if err = dfr.GetContext("Start_Booking_Cab", &p); err != nil {
	// 	log.Println("Bad request from webhook - Param error")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// Do things with the context you just retrieved
	dff := &Fulfillment{
		FulfillmentMessages: Messages{
			ForGoogle(SingleSimpleResponse("hello", "hello")),
			{RichMessage: Text{Text: []string{"hello"}}},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dff)
}
