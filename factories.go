package main

//CityFactory creates a City
func CityFactory(name string) City {
	return City{Name: name, Channel: make(chan ArrivedEvent), Roads: make(Roads)}
}

//MonsterFactory for testing
func MonsterFactory(name string) Monster {
	return Monster{Name: name}
}
