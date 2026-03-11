package main

import (
	"encoding/json"
	"fmt"
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
	//"checkknife0":    10000,

	"truedragonslayer": 4000,
	"voidstaff":        24750,
	"sacrifice":        1450,
	"rbmarble":         80,

	"checkknife0":       883,
	"sporkofdoom":       11000,
	"lightningkarambit": 2400,
	"rumbit":            2400,

	"visible":               2400,
	"soulscendo":            2500,
	"voidscendo":            1250,
	"skeletonknifeamethyst": 1200,

	"skeletonkniferuby": 1200,
	"realitybusty":      1200,
	"epickatana":        1200,
	"galaxyknife":       1200,

	"alienknife":   1200,
	"checkknife6":  224,
	"banhammer":    36,
	"karambitdark": 120,

	"karambit":      120,
	"checkknife5":   120,
	"scissors":      80,
	"horrormachete": 80,

	"flowermachete":         80,
	"skeletonknifeobsidian": 80,
	"null":                  80,
	"eruptionknife":         80,

	"catknife":   80,
	"chaosknife": 18,

	"eyelander":  40000,
	"hyperbeam:": 1500,
	"rapidfire":  35000,
	//"soulscendo": 4000,

	"c4":          1249,
	"awp2":        75000,
	"tacticooler": 2000,
	"bananbit":    12000,

	//"horrormachete": 125000,
	"deflectperk": 150000,
	"fakec4":      1249,
	"tacticool":   300,

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

func HandleGetRequest(response_writer http.ResponseWriter, request_info *http.Request) {
	response_headers := request_info.Header
	access_token := response_headers.Get("Authorization")

	//	fmt.Println(access_token, ACCESS_TOKEN)

	if access_token == ACCESS_TOKEN {
		encoded_vals, _ := json.Marshal(item_values)
		_, err := response_writer.Write(encoded_vals)

		if err != nil {
			fmt.Println("Server Error:", err)

			response_writer.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	response_writer.WriteHeader(http.StatusUnauthorized)
}
