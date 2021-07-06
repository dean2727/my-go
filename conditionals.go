heistReady := false
// parentheses are optional!!	
if heistReady {
	fmt.Println("Let's go!")
} else {  // must put the else on the same line as }
fmt.Println("Act normal.")
}
// also have else if (which also must be on same line as })

lockCombo := "2-35-19"
robAttempt := "1-1-1"
// == and != for comparison. also, we have <, >, <=, and >=
if (lockCombo == robAttempt) {
	fmt.Print("The vault is now opened.")
}

// logical operators: &&, ||, and !
rightTime := true
rightPlace := true
if rightTime && rightPlace {
	fmt.Println("We're outta here!")
} else {
	fmt.Println("Be patient...")
}
enoughRobbers := false
enoughBags := true
if enoughRobbers || enoughBags {
	fmt.Println("Grab everything!")
} else {
	fmt.Println("Grab whatever you can!")
}

// switch statement
clothingChoice := "sweater"
switch clothingChoice {
case "shirt":
  fmt.Println("We have shirts in S and M only.")
case "polos":
  fmt.Println("We have polos in M, L, and XL.")
case "sweater":
  fmt.Println("We have sweaters in S, M, L, and XL.")
case "jackets":
  fmt.Println("We have jackets in all sizes.")
default:
  fmt.Println("Sorry, we don't carry that")
}

// scoped short declaration statement. declaring a variable in the condition of if or switch statements
// these variables are scoped to the conditions blocks
x := 8
y := 9
if product := x * y; product > 60 {
  fmt.Println(product, "  is greater than 60")
}
switch season := "summer" ; season {
case "summer":
  fmt.Println("Go out and enjoy the sun!")
}