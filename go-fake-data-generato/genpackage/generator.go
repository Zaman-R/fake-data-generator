package genpackage

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

type Info struct {
	FirstName   string `faker:"firstName,unique"`
	LastName    string `faker:"lastName,unique"`
	PhoneNumber string `faker:"phoneNumber,unique"`
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
		x := Info{}
		err := faker.FakeData(&x)
		if err != nil {
			fmt.Println(err)
		}
		// Generate a random phone number with the specified format
		x.PhoneNumber = generatePhoneNumber()
		infos = append(infos, x)
	}
	return infos
}
