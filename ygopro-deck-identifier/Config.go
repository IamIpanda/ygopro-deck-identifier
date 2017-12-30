package ygopro_deck_identifier

import (
	"os"
	"encoding/json"
)

type Configuration struct {
	DeckDefPath string
	UnknownDeck string
	IdentifierNames []string
	Listening string
	AccessKey string
}

var Config Configuration

func InitializeConfig() {
	file, err := os.Open("./ygopro-deck-identifier/Config.json")
	if err != nil {
		Logger.Errorf("Failed to open Config.json. %v", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		Logger.Errorf("Failed to load config: %v", err)
	}
}