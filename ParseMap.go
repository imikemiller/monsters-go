package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

// ParseMap accepts the filePath of a map text file and parses it into an array of City
func ParseMap(path string) Cities {
	//some
	line := []byte{'\n'}
	space := []byte{' '}
	equals := []byte{'='}

	//get the file contents
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	//close when the func returns
	defer file.Close()

	//m contains the byte array
	m, err := ioutil.ReadAll(file)
	//split by newline
	//each line represents a city
	cities := bytes.Split(m, line)

	//init some vars
	parsedMap := make(Cities)
	roadList := make(map[string]map[string]string)
	for _, c := range cities {
		if len(c) > 0 {

			cityNameAndRoads := bytes.Split(c, space)
			name := string(cityNameAndRoads[0])
			//add the city to the map indexed by its name
			parsedMap[name] = CityFactory(name)

			//init inner maps
			roadList[name] = make(map[string]string)

			//start at index 1 to ignore the city name item
			for i := 1; i < len(cityNameAndRoads); i++ {
				directionAndDestination := bytes.Split(cityNameAndRoads[i], equals)
				direction := string(directionAndDestination[0])
				destination := string(directionAndDestination[1])
				roadList[name][direction] = destination
			}
		}
	}

	//inject the roads for each city creating a Road with a reference to destination to the City
	for city, roads := range roadList {
		for direction, destination := range roads {
			parsedMap[city].Roads[destination] = Road{Direction: direction, Destination: parsedMap[destination]}
		}
	}

	return parsedMap
}
