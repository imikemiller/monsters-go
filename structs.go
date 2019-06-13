package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"
)

//Road has a direction and a destination
type Road struct {
	Direction   string
	Destination City
}

//Roads container
type Roads map[string]Road

//Random ly return a road from the container
func (roads Roads) Random() Road {
	keys := reflect.ValueOf(roads).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}

	keysLen := len(strkeys)
	if keysLen <= 0 {
		log.Fatal("No keys")
	}
	randKey := strkeys[rand.Intn(keysLen)]
	return roads[randKey]

}

//City has a name, upto 4 roads (N,S,E,W) and has 0 to many monsters visiting
type City struct {
	Name    string
	Roads   Roads
	Channel chan ArrivedEvent
}

//Arrival - fired by City when a Monster arrives
func (city City) Arrival(monster Monster) {
	go func(channel chan ArrivedEvent) {
		channel <- ArrivedEvent{Monster: monster, City: city}
	}(city.Channel)
}

//Cities container
type Cities map[string]City

//Random ly return a city from the container
func (cities Cities) Random() City {
	keys := reflect.ValueOf(cities).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	keysLen := len(strkeys)
	if keysLen <= 0 {
		log.Fatal("No keys")
	}
	randKey := strkeys[rand.Intn(keysLen)]
	return cities[randKey]
}

//Monster has a name, and city (location) and a number of days travelled
type Monster struct {
	Name    string
	Days    int
	City    City
	Dead    chan bool
	Arrived chan City
}

//Roam goroutine function
func (monster Monster) Roam(cities Cities) {
	//find a starting city and arrive
	monster.Arriving(cities.Random())

	go func() {
		for {
			city := <-monster.Arrived
			func(city City) {
				fmt.Println("Roamed to city: ", city.Name)
				for monster.Days < 10000 {
					select {
					//If a dead message is sent when the monster dies
					case <-monster.Dead:
						return
					default:
						//Make a move
						time.Sleep(1 * time.Millisecond)
						monster.Move()
					}
				}
			}(city)
		}
	}()
}

//Move down a random road from the current city
func (monster Monster) Move() {
	if len(monster.City.Roads) == 0 {
		log.Fatal("No city", monster)
	}
	road := monster.City.Roads.Random()
	monster.Arriving(road.Destination)
}

//Arriving to a city
func (monster Monster) Arriving(city City) {
	monster.Days++
	monster.City = city
	monster.Arrived <- city
	city.Arrival(monster)
}

//Monsters container
type Monsters []Monster

//Shuffle the indexes of the monsters
func (monsters Monsters) Shuffle() {

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(monsters) > 0 {
		n := len(monsters)
		randIndex := r.Intn(n)
		monsters[n-1], monsters[randIndex] = monsters[randIndex], monsters[n-1]
		monsters = monsters[:n-1]
	}
}

//Take first n items
func (monsters Monsters) Take(totalMonsters int) Monsters {
	return monsters[0:totalMonsters]
}

//ArrivedEvent is a container for the info related to a monsters arrival in a City
type ArrivedEvent struct {
	Monster Monster
	City    City
}
