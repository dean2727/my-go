package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }

  openedVault := rand.Intn(100)
  if isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("Vault cant be opened!")
  }

  leftSafely := rand.Intn(5)
  if isHeistOn {
    switch leftSafely {
    case 1:
      isHeistOn = false
      fmt.Println("Heist failed! Caught by a guard!")
    case 2:
      isHeistOn = false
      fmt.Println("Heist failed! Tripped and died!")
    case 3:
      isHeistOn = false
      fmt.Println("Heist failed! You hit a booby trap!")
    case 4:
      isHeistOn = false
      fmt.Println("Heist failed! Eaten by the crocodile!")
    default:
      fmt.Println("Start the getaway car!")
    }

    if leftSafely != 0 {
      amtStolen := 10000 + rand.Intn(1000000)
      fmt.Println("Successfully stole %..2f!", amtStolen)
    }
  }

  fmt.Printf("Heist on: %v", isHeistOn)
}