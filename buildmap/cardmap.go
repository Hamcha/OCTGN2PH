package main

import (
	"log"
)

// GUID to Card ID map
type cardMap map[string]string

func buildMap(sets []xmlSet) cardMap {
	cdb := make(cardMap)
	for _, set := range sets {
		log.Printf("Generating card map for \"%s\"...\n", set.Name)
		for _, card := range set.Cards {
			// Find card number
			number := ""
			for _, prop := range card.Properties {
				if prop.Name == "Number" {
					number = prop.Value
					break
				}
			}
			if number == "" {
				continue
			}
			cdb[card.ID] = number
		}
	}
	log.Printf("Generated card map (%d cards total)\n", len(cdb))
	return cdb
}
