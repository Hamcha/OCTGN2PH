package main // import "github.com/Hamcha/OCTGN2PH/o2ph"
import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type xmlDeck struct {
	Cards []xmlCard `xml:"section>card"`
}

type xmlCard struct {
	Qty    int    `xml:"qty,attr"`
	CardID string `xml:"id,attr"`
	Name   string `xml:",chardata"`
}

func main() {
	cardmap := flag.String("cardmap", "cards.json", "Path to cards.json (generated by buildmap)")
	indeck := flag.String("deck", "deck.o8d", "Path to deck file to convert")
	flag.Parse()

	// Open card map
	mapfile, err := os.Open(*cardmap)
	if err != nil {
		log.Fatalf("Could not open card map %s: %s\n", *cardmap, err.Error())
	}

	// Parse card map
	var cdb map[string]string
	err = json.NewDecoder(mapfile).Decode(&cdb)
	mapfile.Close()
	if err != nil {
		log.Fatalf("Could not parse card map: %s\n", err.Error())
	}

	// Open deck file
	deckfile, err := os.Open(*indeck)
	if err != nil {
		log.Fatalf("Could not open deck file %s: %s\n", *indeck, err.Error())
	}
	defer deckfile.Close()

	// Read deck
	var deck xmlDeck
	err = xml.NewDecoder(deckfile).Decode(&deck)
	if err != nil {
		log.Fatalf("Could not parse deck: %s\n", err.Error())
	}

	// Build URL string
	cardlist := ""
	for _, card := range deck.Cards {
		// Check if card id is in map
		cid, ok := cdb[card.CardID]
		if !ok {
			log.Fatalf("Card ID %s (%s) not found in card map\n", card.Name, card.CardID)
		}
		cardlist += fmt.Sprintf("%sx%d-", strings.ToLower(cid), card.Qty)
	}
	cardlist = strings.TrimRight(cardlist, "-")
	fmt.Printf("http://ponyhead.com/deckbuilder?v1code=%s\n", cardlist)
}