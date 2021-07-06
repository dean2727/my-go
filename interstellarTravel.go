package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("Fuel level:", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
  case "Venus":
    fuel = 300000
  case "Mercury":
    fuel = 500000
  case "Mars":
    fuel = 700000
  default:
    fuel = 0
  }

  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("Welcome, beautiful alien visitor! You are about to enter", planet, ". We hope you enjoy your stay!")
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  }
  if fuelCost > fuelRemaining {
    cantFly()
  }
  return fuelRemaining
}

func main() {
  fuelGauge(5)
  fmt.Println(calculateFuel("Venus"))
  fmt.Println(calculateFuel("Mars"))
  fmt.Println(calculateFuel("Mercury"))
  fmt.Println(calculateFuel("Earth"))
  
  var fuel int
  fuel = 1000000
  planetChoice := "Venus"
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
}