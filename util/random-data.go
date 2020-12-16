package main

// see: https://godoc.org/github.com/Pallinder/go-randomdata

import (
	// "encoding/json"
	"fmt"
 	// "os"

	"github.com/Pallinder/go-randomdata"
)

func main() {
   fmt.Println(randomdata.SillyName())
   fmt.Println(randomdata.FirstName(randomdata.RandomGender))
   fmt.Println(randomdata.LastName())
   fmt.Println(randomdata.FullName(randomdata.RandomGender))
   fmt.Println(randomdata.Email())
   fmt.Println(randomdata.Address())
   fmt.Println(randomdata.Number(100, 1000))
   fmt.Println(randomdata.StringNumber(2, "-"))
   fmt.Println(randomdata.IpV4Address())
   fmt.Println(randomdata.IpV6Address())

   fmt.Println(randomdata.Letters(100))
   // fmt.Println(randomdata.RandStringRunes(100))

    /*
	profiles := []interface{}{}
	for i := 0; i < 100; i++ {
		profile := randomdata.GenerateProfile(randomdata.RandomGender)
		// fmt.Printf("%v\n", profile)
		profiles = append(profiles, profile)
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&profiles); err != nil {
		fmt.Println(err)
	}
	*/
}
