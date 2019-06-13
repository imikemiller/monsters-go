package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestInitialiseMonsters(t *testing.T) {
	monsters := initialiseMonsters(10)

	set1 := monsters.Take(10)
	for i := 0; i < len(set1); i++ {
		fmt.Println(set1[i].Name)
	}

	monsters.Shuffle()

	set2 := monsters.Take(10)
	for i := 0; i < len(set2); i++ {
		fmt.Println(set2[i].Name)
	}
	
	assert.Assert(t, len(set1) == 10)
}
