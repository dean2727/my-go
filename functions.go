package main
 
import "fmt"

// global scope: contains main() and other functions
// local scope: within each function

// super simple function
func summonNicole() {
	fmt.Println("Hey Nicole, get over here!")
}

func getLengthOfCentralPark() int32 {
	var lengthInBlocks int32
	lengthInBlocks = 51
	return lengthInBlocks
}

func multiplier(x int32, y int32) int32 {
	return x * y
}

// multiple return values are possible
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
	averageGrade := (midtermGrade + finalGrade) / 2
	var gradeLetter string
	if averageGrade > 90 {
	  gradeLetter = "A"
	} else if averageGrade > 80 {
	  gradeLetter = "B"
	} else if averageGrade > 70 {
	  gradeLetter = "C"
	} else if averageGrade > 60 {
	  gradeLetter = "D"
	} else {
	  gradeLetter = "F"
	}
	return gradeLetter, averageGrade 
}

// the defer keyword is used before a line of code to say "hey, I want you to execute at the end of the function"
// this is useful for logging, file writing, and other utilities
// it's also helpful because we dont have to find all the return statements and try to place lines of code before each one
func queryDatabase(query string) string {
	var result string
	connectDatabase()
	defer disconnectDatabase()
	if query == "SELECT * FROM coolTable;" {
	  result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
	}  
	fmt.Println(result)
	return result
}
func connectDatabase() {
	fmt.Println("Connecting to the database.")
}
func disconnectDatabase() {
	fmt.Println("Disconnecting from the database.")
}

func main() {
	summonNicole()

	var centralParkLength int32
	centralParkLength = getLengthOfCentralPark()
	fmt.Println(centralParkLength) // Prints: 51

	var product int32
	product = multiplier(25, 4)
	fmt.Println(product) // Prints: 100

	var myMidterm, myFinal float32
	myMidterm = 89.4
	myFinal = 74.9
	var myAverage float32
	var myGrade string
	myGrade, myAverage = GPA(myMidterm, myFinal)
	fmt.Println(myAverage, myGrade) // Prints 82.12 B

	queryDatabase("SELECT * FROM coolTable;")
}
