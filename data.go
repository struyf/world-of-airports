package data

// Represents a basic airport for your game
type Airport struct {
    Name       string 
    Grid       [][]int      // Placeholder for representing the airport layout 
    Runways    []Runway     // Could define a more detailed Runway struct later
    Funds      int               
} 
