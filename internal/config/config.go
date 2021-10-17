package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	CDPURL                string   `json:"cdpurl"`
	NotificationProviders []string `json:"notification_providers"`
	GotifyURL             string   `json:"gotify_url"`
}

var Settings Config

func Load() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln("CONFIG.LOAD", err)
	}
	err = json.Unmarshal(data, &Settings)
	if err != nil {
		log.Fatalln("CONFIG.LOAD unmarshall", err)
	}
}
