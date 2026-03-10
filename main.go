package main

import (
	//"fmt"
	"encoding/json"
	"log"
	"net/http"
	"os"
	// "strings"
)

var ACCESS_TOKEN string = "e7b8c1f3-8c1d-4f3e-a1d2-2f9b5e8c1234"

var item_values = map[string]int{
	"angelperk":      900,
	"truer6revolver": 10000,
	"r6revolver":     5000,
	"checkknife0":    10000,

	"eyelander":  40000,
	"hyperbeam:": 1500,
	"rapidfire":  35000,
	"soulscendo": 4000,

	"c4":          1249,
	"awp2":        75000,
	"tacticooler": 2000,
	"bananbit":    12000,

	"horrormachete": 125000,
	"deflectperk":   150000,
	"fakec4":        1249,
	"tacticool":     300,

	"bananagun":   30000,
	"lostandsoul": 2000,
	"candy":       20000000,
	"tomahawk":    6000,
}

func main() {
	http.HandleFunc("/itemvalues", HandleGetRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HandleGetRequest(response_writer http.ResponseWriter, _ *http.Request) {
	response_headers := response_writer.Header()
	access_token := response_headers.Get("Authorization")

	if access_token == ACCESS_TOKEN {
		encoded_vals, _ := json.Marshal(item_values)
		response_writer.Write(encoded_vals)
	}
}
