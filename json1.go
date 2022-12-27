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
	Height float64
}

func main() {
	
	// person := Person{
	// 	"kim", 20, "Uganda",
	// }
	person2 := Person{
		Name:"kim",Age: 47,Location:"Mbarara",Height:788,
	}

	personArray, err := json.Marshal(person2)

	if err != nil{
		log.Fatalf("Unable to Encode")
	}

	fmt.Println(string(personArray))

}
