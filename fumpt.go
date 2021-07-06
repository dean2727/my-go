package main

// fumpt package: for printing and formatting data
import "fmt"

// Println(): prints its arguments with a space between each argument and a newline at the end
fmt.Println("Println", "formats", "really well")

// Print(): prints its args with no spaces and no newline
fmt.Print("The answer is", ": ")
fmt.Print("12")
// prints "The answer is: 12"

// Printf(): allows us to use string placeholders using verbs, formatted %<letter> (no newline included)
animal1 := "cat"
animal2 := "dog"
fmt.Printf("Are you a %v or a %v person?", animal1, animal2)
specialNum := 42
// %T: tells us the type
fmt.Printf("This value's type is %T.", specialNum)
// %<decimal>f: for number precision
gpa := 3.8
fmt.Printf("You're averaging: %.2f.", gpa)
// Prints: You're averaging 3.80.

// Sprint() and Sprintln() are like stringbuilders (Sprintln auto inputs the spaces and newline)
teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)
fmt.Print(teacherSays)
// Prints: You scored a 100 on the test! Great job!
quote = fmt.Sprintln("Look ma,", "no spaces!")
fmt.Print(quote) // Prints Look ma, no spaces!
// Sprintf(): allows interpolation, similar to Printf()
correctAns := "A"
answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)
fmt.Print(answer) // Prints: And the correct answer is… A!

// Scan(): user input. takes an address for a variable as an argument (&arg), and is similar to cin in c++
