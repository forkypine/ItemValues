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
	// Crates
	"crateepic":       12,
	"crateRrare":      5,
	"craterare":       2,
	"crateRcommon":    3,
	"cratecommon":     1,
	"cratestock":      0,
	"crateRstock":     0,
	"crateRepic":      15,
	"crateRlegendary": 20,
	"crateRmythical":  35,
	"diamondhead": 1000000,
	"angelperk": 1000000,

	// Contrabands

	"annihilator":       16500,
	"hellsannihilator":  16500,
	"longshot":          120,
	"tangledr6revolver": 14480,

	"truer6revolver":    14400,
	"r6revolver":        11000,
	"encodedr6revolver": 11000,

	"truedragonslayer": 4000,
	"voidstaff":        16500,
	"sacrifice":        16500,
	"rbmarble":         80,

	"checkknife0": 883,
	"sporkofdoom": 11000,
	"myriadblade": 11000,
	"eyelander":   40000,

	// Mythicals

	"amethystrevolver": 2000,
	"gemrevolver":      1000,
	"spadebreaker":     2000,
	"toolgun":          2000,

	"phantomawp": 1800,
	"freezeawp":  1800,
	"rapidfire":  400,
	"deathawp":   1200,

	"neoawp": 1200,

	"lightningkarambit": 1800,
	"rumbit":            1800,

	"visibleknife":          1800,
	"soulscendo":            2500,
	"voidscendo":            1250,
	"skeletonknifeamethyst": 1200,

	"skeletonkniferuby": 1200,
	"realitybuster":      1200,
	"epickatana":        1200,
	"galaxyknife":       1200,

	"alienknife":   1200,
	"checkknife6":  224,
	"banhammer":    36,
	"karambitdark": 120,

	// Legendaries

	"awplegacy":   180,
	"awp":         180,
	"greedandsin": 40,
	"tacticooler": 88,

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

	// Uniques

	"usp-s":           45,
	"sixpistols":      45,
	"madpaintballgun": 45,
	"greedrevolver":   20,

	"sinrevolver":       20,
	"usp-45":            20,
	"lunacyandsolstice": 10,
	"lostandsoul":       10,

	"shotgunrevolver": 34,
	"goldluger":       20,
	"tacticool":       8,

	"spear":         45,
	"pianobreaker":  20,
	"technology":    20,
	"skeletonknife": 20,

	"piecemaker":        20,
	"skylark":           20,
	"flambe":            20,
	"heartbreakerknife": 20,

	"mm2deathshard": 20,
	"darkknife":     20,
	"midnightknife": 20,
	"dragonlore":    20,

	"robuxknife":    20,
	"tixknife":      20,
	"tomahawkknife": 20,
	"stingerknife":  20,

	"overseerknife2":   20,
	"crescendo":        8,
	"goldwrench":       8,
	"dearestgratitude": 8,

	"checkknife4": 56,
	"checkknife3": 28,

	// Epics

	"solsticerevolver": 5,
	"lunacyrevolver":   5,
	"lostrevolver":     5,
	"soulrevolver":     5,

	"voidrevolver": 5,
	"corruptluger": 5,

	"checkknife2": 14,
	"checkknife":  7,

	// "hyperbeam:": 1500,
	//"soulscendo": 4000,

	// "c4":       1249,
	// "bananbit": 12000,

	//"horrormachete": 125000,
	// "deflectperk": 150000,
	// "fakec4":      1249,

	// Commons

	"desperado":          2,
	"rosequartzrevolver": 1,

	// MISC

	"spamknife":   1200,
	"customtesla": 360,
	"custombeam":  240,
	"customscope": 160,

	"teslaknife":        68,
	"shotgunknife":      30,
	"mvppassvoucher_bm": 130,
	"vippassvoucher_bm": 60,

	"crateperi_bm":  120,
	"crateprem_bm":  15,
	"crateprem2_bm": 15,
	"cratet1_bm":    6,

	"hyperbeam":      60,
	"deathbeam":      24,
	"noscopeknife":   2,
	"mvppassvoucher": 18,

	"nametag":   18,
	"cratet1_f": 1,
	"cheapkarambit": 5,
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

	if request_info.Header.Get("Content-Type") != "application/json" {
		response_writer.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

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
