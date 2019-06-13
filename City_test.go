package main

import (
	"fmt"
	"testing"
)

func TestCity(t *testing.T) {
	city := CityFactory("this_string")
	finished := make(chan bool)

	fmt.Println(city)

	go func(finished chan bool) {
		for {
			event := <-city.Channel
			func(event ArrivedEvent) {
				fmt.Println(event.City.Name)
				finished <- true
			}(event)
		}
	}(finished)

	monster := MonsterFactory("baddie")

	city.Arrival(monster)
	<-finished
}
