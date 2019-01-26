package main // import "github.com/Hamcha/OCTGN2PH/buildmap"

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func main() {
	doc := getDocumentDir()
	defaultSetPath := filepath.Join(doc, "OCTGN", "GameDatabase", "65656467-b709-43b2-a5c6-80c2f216adf9", "Sets")
	setpath := flag.String("set-path", defaultSetPath, "Path to sets")
	flag.Parse()

	// Load sets
	sets, err := loadSets(*setpath)
	if err != nil {
		log.Fatalf("Could not load all sets: %s\nClosing...\n", err.Error())
	}

	cmap := buildMap(sets)
	err = json.NewEncoder(os.Stdout).Encode(cmap)
	if err != nil {
		log.Fatalf("Error while encoding to JSON: %s\n", err.Error())
	}
}

func getDocumentDir() string {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\User Shell Folders`, registry.QUERY_VALUE)
	if err != nil {
		return ""
	}
	defer k.Close()

	s, _, err := k.GetStringValue("Personal")
	if err != nil {
		return ""
	}
	return s
}
