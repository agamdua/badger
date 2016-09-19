package main

import (
	"github.com/asdine/storm"
	"log"
)

type OutboundConfig struct {
	Action       string `storm:"id"`
	URL          string `storm:"unique"`
	Verb         string
	APIKeyHeader string `storm:"unique"`
	APIKey       string `storm:"unique"` // TODO: enhance for headers
}

func GetConfig(message Message) OutboundConfig {
	db, err := storm.Open("my.db")

	if err != nil {
		log.Fatal(err)
	}

	var config OutboundConfig

	// TODO: remove test data
	testConfig := OutboundConfig{
		Action:       "getPictureData",
		URL:          "https://api.projectoxford.ai/vision/v1.0/models",
		Verb:         "GET",
		APIKeyHeader: "Ocp-Apim-Subscription-Key",
		APIKey:       getAPIKey(),
	}

	saveError := db.Save(&testConfig)

	if saveError != nil {
		log.Fatal(err)
	}

	queryError := db.One("Action", message.Action, &config)
	log.Print(config)

	defer db.Close()

	if queryError != nil {
		log.Fatal(err)
	}

	return config
}
