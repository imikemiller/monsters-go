package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestParseMap(t *testing.T) {

	parsedMap := ParseMap("maps/world_map_small.txt")

	for _, city := range parsedMap {
		fmt.Println("city:", city.Name, "->")
		for _, road := range city.Roads {
			fmt.Println(road.Direction, "goes to", road.Destination.Name)
		}
		fmt.Println("end city")
	}

	assert.Assert(t, is.Len(parsedMap, 28))

}
