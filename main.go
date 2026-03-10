package main

import (
	// "fmt"
	"log"
	"net/http"

	"encoding/json"
	// "strings"
)

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

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func HandleGetRequest(response_writer http.ResponseWriter, _ *http.Request) {
	encoded_vals, _ := json.Marshal(item_values)
	response_writer.Write(encoded_vals)
}
