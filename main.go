package main

import (
	"fmt" 
	"math/rand"
	"time"
	"strconv"
)

type Date struct {
	month, date int
}

type Flight struct {
	id, origin, destination string
	price                   int
	departure               Date
}

type tabFlight [1000000]Flight

var TF tabFlight
var nFlight int

func main() {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())

	var n int
	n = 100
	inputFlightData(n)
	TF[n-1].origin = "B"
	TF[n-1].destination = "C"

	var choice string
	choice = "2"
	if choice == "1" {
		searchFlight("B", "C")	
	} else if choice == "2" {
		searchFlightRecursive("B", "C", 0)
	} else {
		fmt.Println("Invalid choice")
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %.5f seconds\n", elapsed.Seconds())
}

func inputFlightData(n int) {
	for i := 0; i < n; i++ {
		randomNumber := rand.Intn(1001)
		randomString := strconv.Itoa(randomNumber)
		TF[i].id = randomString
		TF[i].origin = "A"
		TF[i].destination = "B"
		TF[i].price = 123
		nFlight++
	}
}

func searchFlight(origin, destination string) {
	var found bool
	for i := 0; i < nFlight; i++ {
		if TF[i].origin == origin && TF[i].destination == destination {
			found = true
		}
	}
	if found {
		fmt.Println("Flight found")
	} else {
		fmt.Println("Flight not found")
	}
}

func searchFlightRecursive(origin, destination string, i int) {
	if i == nFlight {
		fmt.Println("Flight not found")
	} else {
		if TF[i].origin == origin && TF[i].destination == destination {
			fmt.Println("Flight found")
		} else {
			searchFlightRecursive(origin, destination, i+1)
		}
	}
}