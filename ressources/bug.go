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

// bug structure
type Bugs map[string]Bug

type Bug struct {
	ID             int64           `json:"id"`
	FileName       string          `json:"file-name"`
	Name           NameBug         `json:"name"`
	Availability   AvailabilityBug `json:"availability"`
	Price          int64           `json:"price"`
	PriceFlick     int64           `json:"price-flick"`
	CatchPhrase    string          `json:"catch-phrase"`
	MuseumPhrase   string          `json:"museum-phrase"`
	ImageURI       string          `json:"image_uri"`
	IconURI        string          `json:"icon_uri"`
	AltCatchPhrase []string        `json:"alt-catch-phrase,omitempty"`
}

type AvailabilityBug struct {
	MonthNorthern      string    `json:"month-northern"`
	MonthSouthern      string    `json:"month-southern"`
	Time               string    `json:"time"`
	IsAllDay           bool      `json:"isAllDay"`
	IsAllYear          bool      `json:"isAllYear"`
	Location           string    `json:"location"`
	Rarity             RarityBug `json:"rarity"`
	MonthArrayNorthern []int64   `json:"month-array-northern"`
	MonthArraySouthern []int64   `json:"month-array-southern"`
	TimeArray          []int64   `json:"time-array"`
}

type NameBug struct {
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

type RarityBug string

const (
	CommonBug    RarityBug = "Common"
	RareBug      RarityBug = "Rare"
	UltraRareBug RarityBug = "Ultra-rare"
	UncommonBug  RarityBug = "Uncommon"
)

// routers
func getOneBugById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Trying to open bug.json...")

	jsonFile, err := os.Open("apidata/bug.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}

	fmt.Println("Successfully opened bug.json!")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var bugs Bugs

	err = json.Unmarshal(byteValue, &bugs)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}

	bugID, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 34)

	if bugID > int64(len(bugs)) {
		fmt.Fprintf(w, "This fish doesn't exist, we have %d bugs in total.", len(bugs))
	} else {
		for _, singleBug := range bugs {
			if singleBug.ID == bugID {
				json.NewEncoder(w).Encode(singleBug)
			}
		}
	}

}
