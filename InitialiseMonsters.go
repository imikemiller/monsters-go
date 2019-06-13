package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func initialiseMonsters(totalMonsters int) Monsters {

	names, err := os.Open("data/names.json")
	defer names.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	//unmarshall the JSON into Monsters array
	byteValue, _ := ioutil.ReadAll(names)
	var monsters Monsters
	json.Unmarshal(byteValue, &monsters)

	for _, monster := range monsters {
		monster.Dead = make(chan bool)
	}
	return monsters.Take(totalMonsters)
}
