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

// sea creature structure
type SeaCreatures map[string]SeaCreature

type SeaCreature struct {
	ID           int64                   `json:"id"`
	FileName     string                  `json:"file-name"`
	Name         NameSeaCreature         `json:"name"`
	Availability AvailabilitySeaCreature `json:"availability"`
	Speed        SpeedSeaCreature        `json:"speed"`
	Shadow       ShadowSeaCreature       `json:"shadow"`
	Price        int64                   `json:"price"`
	CatchPhrase  string                  `json:"catch-phrase"`
	ImageURI     string                  `json:"image_uri"`
	IconURI      string                  `json:"icon_uri"`
	MuseumPhrase string                  `json:"museum-phrase"`
}

type AvailabilitySeaCreature struct {
	MonthNorthern      string            `json:"month-northern"`
	MonthSouthern      string            `json:"month-southern"`
	Time               Time              `json:"time"`
	IsAllDay           bool              `json:"isAllDay"`
	IsAllYear          bool              `json:"isAllYear"`
	Rarity             RaritySeaCreature `json:"rarity"`
	MonthArrayNorthern []int64           `json:"month-array-northern"`
	MonthArraySouthern []int64           `json:"month-array-southern"`
	TimeArray          []int64           `json:"time-array"`
}

type NameSeaCreature struct {
	NameUSen string `json:"name-USen"`
	NameEUen string `json:"name-EUen"`
	NameEUde string `json:"name-EUde"`
	NameEUnl string `json:"name-EUnl"`
	NameEUes string `json:"name-EUes"`
	NameUSes string `json:"name-USes"`
	NameEUfr string `json:"name-EUfr"`
	NameUSfr string `json:"name-USfr"`
	NameEUit string `json:"name-EUit"`
	NameCNzh string `json:"name-CNzh"`
	NameTWzh string `json:"name-TWzh"`
	NameJPja string `json:"name-JPja"`
	NameKRko string `json:"name-KRko"`
	NameEUru string `json:"name-EUru"`
}

type RaritySeaCreature string

const (
	CommonSeaCreature    RaritySeaCreature = "Common"
	RareSeaCreature      RaritySeaCreature = "Rare"
	UltraRareSeaCreature RaritySeaCreature = "Ultra-rare"
)

type Time string

const (
	EmptySeaCreature           Time = ""
	The4Am9PmSeaCreature       Time = "4am - 9pm"
	The4Pm9AmSeaCreature       Time = "4pm - 9am"
	The9Am4Pm9Pm4AmSeaCreature Time = "9am - 4pm & 9pm - 4am"
	The9Pm4AmSeaCreature       Time = "9pm - 4am"
)

type ShadowSeaCreature string

const (
	LargeSeaCreature        ShadowSeaCreature = "Large"
	LargestSeaCreature      ShadowSeaCreature = "Largest"
	ShadowMediumSeaCreature ShadowSeaCreature = "Medium"
	SmallSeaCreature        ShadowSeaCreature = "Small"
	SmallestSeaCreature     ShadowSeaCreature = "Smallest"
)

type SpeedSeaCreature string

const (
	FastSeaCreature        SpeedSeaCreature = "Fast"
	SlowSeaCreature        SpeedSeaCreature = "Slow"
	SpeedMediumSeaCreature SpeedSeaCreature = "Medium"
	StationarySeaCreature  SpeedSeaCreature = "Stationary"
	VeryFastSeaCreature    SpeedSeaCreature = "Very fast"
	VerySlowSeaCreature    SpeedSeaCreature = "Very slow"
)

// routers
func getOneSeaCreatureById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Trying to open seaCreature.json...")

	jsonFile, err := os.Open("apidata/seaCreature.json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}

	fmt.Println("Successfully opened seaCreature.json!")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var seaCreatures SeaCreatures

	err = json.Unmarshal(byteValue, &seaCreatures)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}

	fishID, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 34)

	if fishID > int64(len(seaCreatures)) {
		fmt.Fprintf(w, "This fish doesn't exist, we have %d sea creatures in total.", len(seaCreatures))
	} else {
		for _, singleFish := range seaCreatures {
			if singleFish.ID == fishID {
				json.NewEncoder(w).Encode(singleFish)
			}
		}
	}

}
