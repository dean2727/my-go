import (
	"math/rand"
	"fmt"
	"time"
)
   
func main() {
	// random integer from 0 to 99 (inclusive)
	// will always return the same number, however, so we must seed (starting point for Go to generate random numbers)
	fmt.Println(rand.Intn(100))
	// seed value is default 1. give a new seed value by passing in an integer to rand.Seed()
	// to have truly random behavior, you must pass in a unique number to Seed() each time you run the program
	// 1 way to do this is to use the "time" package
	rand.Seed(time.Now().UnixNano())  // # of nanoseconds since Jan 1st, 1970 (UTC)
  	fmt.Println(rand.Intn(100))
}