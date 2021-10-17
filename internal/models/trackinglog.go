package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TrackingDetail struct {
	Description string `json:"description"`
	Time        string `json:"time"`
	Status      string `json:"status"`
}

type TrackingLog struct {
	TrackingDetails map[string]TrackingDetail `json:"tracking_details"`
	Code            string                    `json:"code"`
	Service         string                    `json:"service"`
}

func (tl *TrackingLog) Load() error {
	filename := "data/" + tl.Code + ".json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		tl.TrackingDetails = make(map[string]TrackingDetail)
		return err
	}
	json.Unmarshal(data, &tl)
	if err != nil {
		log.Println(err)
		tl.TrackingDetails = make(map[string]TrackingDetail)
		return err
	}
	return nil
}

func (tl *TrackingLog) Save() error {
	filename := "data/" + tl.Code + ".json"
	data, err := json.Marshal(tl)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Println(err)
	}
	return nil
}
