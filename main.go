package main

import (
	"fmt"
	"log"
	"errors"
)

type Truck struct {
	id string
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	return errors.New("Some error")
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s is arrived.\n", truck.id)

		// err is defined as a local variable, which can still be accessed down there in the function, so if error needs to be handled further down, then use this.
		// err := processTruck(truck)
		// if err != nil {
		// 	log.Fatalf("Error processing truck: %s", err)
		// }

		// Benefit here is that err will not be available outside the if statement, which means Garbage collector will pick it and clean it up
		if err := processTruck(truck); err != nil {
			log.Fatalf("Error processing truck: %s", err)
		}
	}
}