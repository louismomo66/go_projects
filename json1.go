package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name     string `json:"firstName"`
	Age      float64
	Location string `json:"location,omitempty"`
}

func main() {
	
	// person := Person{
	// 	"kim", 20, "Uganda",
	// }
	person2 := Person{
		// Name:"kimera",Age: 28,
	  Name:"kim",Age: 27.3,
	}

	personArray, err := json.Marshal(person2)

	if err != nil{
		log.Fatalf("Unable to Encode")
	}

	fmt.Println(string(personArray))

}