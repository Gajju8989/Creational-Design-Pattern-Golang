package main

import (
	"design_patterns/builder"
	"fmt"
)

func main() {
	builder1 := builder.NewConcreteHouseBuilder()

	basicHouse := builder1.SetWindows(4).SetDoors(2).SetRooms(3).Build()
	fmt.Println("Basic House: %+v\n", basicHouse)

	// Using a director
	director := builder.NewHouseDirector(builder1)
	luxuryHouse := director.ConstructLuxuryHouse()
	basicHouse = director.ConstructBasicHouse()
	fmt.Println("Luxury House: %+v\n", luxuryHouse)
	fmt.Println("Basic House: %+v\n", basicHouse)
}
