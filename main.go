package main

import (
	"fmt"
	"sync"
)

//Monster arrives

func main() {
	var wg sync.WaitGroup
	parsedMap := ParseMap("maps/world_map_small.txt")
	monsters := initialiseMonsters(10)

	for _, monster := range monsters {
		wg.Add(1)
		monster.Roam(parsedMap)
	}

	for _, city := range parsedMap {
		for {
			event := <-city.Channel
			func(event ArrivedEvent) {
				fmt.Println(event.City.Name)
			}(event)
		}
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}
