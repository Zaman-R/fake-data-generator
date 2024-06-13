package genpackage

import (
	"fmt"
	"math/rand"
	"time"
)

type Info struct {
	FirstName   string `faker:"firstName,unique"`
	LastName    string `faker:"lastName,unique"`
	PhoneNumber string `faker:"phoneNumber,unique"`
}

var FirstNameList = []string{
	"Alice", "Bob", "Charlie", "Diana", "Edward",
	"Fiona", "George", "Hannah", "Ian", "Julia",
}

var LastNameList = []string{
	"Smith", "Johnson", "Williams", "Jones", "Brown",
	"Davis", "Miller", "Wilson", "Moore", "Taylor",
}

//generateFirstName generate a random firstname
func generateFirstName() string {
	return FirstNameList[rand.Intn(len(FirstNameList))]
}

//generateFirstName generate a random lastname
func generateLastName() string {
	return LastNameList[rand.Intn(len(LastNameList))]
}

// generatePhoneNumber generates a random phone number with the specified format
func generatePhoneNumber() string {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Choose a random prefix from the set {"017", "016", "018", "019"}
	prefixes := []string{"017", "016", "018", "019"}
	prefix := prefixes[rand.Intn(len(prefixes))]

	// Generate 9 random digits for the rest of the phone number
	randomPart := rand.Intn(1e9)
	randomDigits := fmt.Sprintf("%09d", randomPart)

	return prefix + randomDigits
}

// GenerateInfo generates n Info structs with fake data
func GenerateInfo(n int) []Info {
	var infos []Info
	for i := 0; i < n; i++ {
		x := Info{
			FirstName:   generateFirstName(),
			LastName:    generateLastName(),
			PhoneNumber: generatePhoneNumber(),
		}
		infos = append(infos, x)
	}
	return infos
}
