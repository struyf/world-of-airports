package main

import (
    "fmt"
    "github.com/struyf/world-of-airports/data"
    "github.com/struyf/world-of-airports/airport_operations"
)

func main() {
    myAirport := airport_operations.CreateAirport("A Nice Airport") // Example usage
    fmt.Println("Welcome to", myAirport.Name)
}

