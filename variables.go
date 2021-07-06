package main

import "fmt"

func main() {
  var stationName string
  var nextTrainTime int8
  var isUptownTrain bool
  /* other types:
	- float
	- float64
	- complex (imaginary numbers, e.g. 3i or 7+2i)
	- complex64
	- int
	- int16
	- int32
	- int64
	- uint8
	- uint16
	- uint32
	- uint64
	(15 ways to describe a number in Go! WOW!)
  */
  
  stationName = "Union Square"  // must use double quotes
  nextTrainTime = 12
  isUptownTrain = false
  
  fmt.Println("Current station:", stationName)  // comma adds a space like in python
  fmt.Println("Next train:", nextTrainTime, "minutes")  // dont need to convert to string (we love that)
  fmt.Println("Is uptown:", isUptownTrain)
  
  stationName = "Grand Central"
  nextTrainTime = 3
  isUptownTrain = true
  
  fmt.Println("")
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)

  // constants
  const earthsGravity = 9.80665

  // error will be thrown, because variable is declared but not used
  // if we use it in any way, e.g. in fmt.Println(), then no error
  var numberWheels int8

  // string concatenation (+)
  var favoriteSnack string = "baked layes"
  fmt.Println("My favorite snack is " + favoriteSnack)

  /* default values on undefined variables:
	- int: 0
	- string: ""
	- bool: false
  */

  // short declaration operator (:=): if we know what value we want our variable to store on creation
  // no need to specify type, Go infers it
  nuclearMeltdownOccurring := true
  radiumInGroundWater := 4.521
  consolationPrizes := 2  // type: int
  // same as
  var nuclearMeltdownOccurring = true
  var radiumInGroundWater = 4.521
  var consolationPrizes = 2
  
  // int and uint will be either 32 or 64 bit, depending on computer architecture
  var timesWeWereFooled int
  var foolishGamesPlayed uint

  // the operators +=, -=, *=, /= all work in Go

  // defining multiple variables on one line
  var part1, part2 string
  part1 = "To be..."
  part2 = "Not to be..."
  quote, fact := "Bears, Beets, Battlestar Galactica", true
}