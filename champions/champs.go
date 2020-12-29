package champions

import (
	"log"
)

// Champs func
func Champs(fileName string) []Champion {

	champions, err := LoadChampions(fileName)
	if err != nil {
		log.Fatalf("ooops something went wrong when parsing the data, err=%v", err)
	}

	bobo := Champion{
		Name:    "Bobo",
		Classes: []string{"Blademaster", "Gunslinger"},
		Origins: []string{"Ninja"},
		Cost:    2,
	}

	newChamp := CreateChampion(bobo)

	champions = append(champions, newChamp)

	return champions

}
