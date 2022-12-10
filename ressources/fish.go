package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// fish structure
type Fishes map[string]Fish

type Fish struct {
	ID             int64        `json:"id"`
	FileName       string       `json:"file-name"`
	Name           Name         `json:"name"`
	Availability   Availability `json:"availability"`
	Shadow         string       `json:"shadow"`
	Price          int64        `json:"price"`
	PriceCj        int64        `json:"price-cj"`
	CatchPhrase    string       `json:"catch-phrase"`
	MuseumPhrase   string       `json:"museum-phrase"`
	ImageURI       string       `json:"image_uri"`
	IconURI        string       `json:"icon_uri"`
	AltCatchPhrase []string     `json:"alt-catch-phrase,omitempty"`
}

type Availability struct {
	MonthNorthern      string   `json:"month-northern"`
	MonthSouthern      string   `json:"month-southern"`
	Time               Time     `json:"time"`
	IsAllDay           bool     `json:"isAllDay"`
	IsAllYear          bool     `json:"isAllYear"`
	Location           Location `json:"location"`
	Rarity             Rarity   `json:"rarity"`
	MonthArrayNorthern []int64  `json:"month-array-northern"`
	MonthArraySouthern []int64  `json:"month-array-southern"`
	TimeArray          []int64  `json:"time-array"`
}

type Name struct {
	NameUSen string `json:"name-USen"`
	NameEUen string `json:"name-EUen"`
	NameEUde string `json:"name-EUde"`
	NameEUes string `json:"name-EUes"`
	NameUSes string `json:"name-USes"`
	NameEUfr string `json:"name-EUfr"`
	NameUSfr string `json:"name-USfr"`
	NameEUit string `json:"name-EUit"`
	NameEUnl string `json:"name-EUnl"`
	NameCNzh string `json:"name-CNzh"`
	NameTWzh string `json:"name-TWzh"`
	NameJPja string `json:"name-JPja"`
	NameKRko string `json:"name-KRko"`
	NameEUru string `json:"name-EUru"`
}

type Location string

const (
	Pier                    Location = "Pier"
	Pond                    Location = "Pond"
	River                   Location = "River"
	RiverClifftop           Location = "River (Clifftop)"
	RiverClifftopPond       Location = "River (Clifftop) & Pond"
	RiverMouth              Location = "River (Mouth)"
	Sea                     Location = "Sea"
	SeaWhenRainingOrSnowing Location = "Sea (when raining or snowing)"
)

type Rarity string

const (
	Common    Rarity = "Common"
	Rare      Rarity = "Rare"
	UltraRare Rarity = "Ultra-rare"
	Uncommon  Rarity = "Uncommon"
)

type Time string

const (
	Empty           Time = ""
	The4Am9Pm       Time = "4am - 9pm"
	The4Pm9Am       Time = "4pm - 9am"
	The9Am4Pm       Time = "9am - 4pm"
	The9Am4Pm9Pm4Am Time = "9am - 4pm & 9pm - 4am"
	The9Pm4Am       Time = "9pm - 4am"
)

// routers
func getOneFishById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Trying to open fish.json...")

	jsonFile, err := os.Open("apidata/fish.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}

	fmt.Println("Successfully opened fish.json!")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var fishes Fishes

	err = json.Unmarshal(byteValue, &fishes)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}

	fishID, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 34)

	if fishID > int64(len(fishes)) {
		fmt.Fprintf(w, "This fish doesn't exist, we have %d fishes in total.", len(fishes))
	} else {
		for _, singleFish := range fishes {
			if singleFish.ID == fishID {
				json.NewEncoder(w).Encode(singleFish)
			}
		}
	}

}
