package main

import (
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
)

type xmlSet struct {
	Name  string    `xml:"name,attr"`
	Cards []xmlCard `xml:"cards>card"`
}

type xmlCard struct {
	ID         string        `xml:"id,attr"`
	Properties []xmlProperty `xml:"property"`
}

type xmlProperty struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func loadSetCards(filename string) (xmlSet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return xmlSet{}, err
	}
	defer file.Close()

	var set xmlSet
	err = xml.NewDecoder(file).Decode(&set)
	return set, err
}

func loadSets(basedir string) ([]xmlSet, error) {
	var sets []xmlSet
	err := filepath.Walk(basedir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error while scanning %s: %s\n", path, err.Error())
			return err
		}
		// Ignore non XML files
		ext := filepath.Ext(path)
		if info.IsDir() || ext != ".xml" {
			return nil
		}
		set, err := loadSetCards(path)
		if err != nil {
			return err
		}
		log.Printf("Loaded set: %s (%d cards)\n", set.Name, len(set.Cards))
		sets = append(sets, set)
		return nil
	})
	return sets, err
}
