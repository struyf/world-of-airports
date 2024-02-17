package airport_operations

import "github.com/struyf/world-of-airports/data"

// A simple function to create a new airport (for now)
func CreateAirport(name string) *data.Airport {
    return &data.Airport{
        Name:  name,
        Funds: 10000, // Starting money
    }
}
