package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {

	jsonStream := `{"Name":"Kim","Age":20,"Location":"Texas"}{"Name":"Jeff","Age":50,"Location":"Wampewo"}`

	reader := strings.NewReader(jsonStream)

	writer:= os.Stdout

	decoder:=json.NewDecoder(reader)

	encoder:= json.NewEncoder(writer)

	for{
		var v map[string]interface{}

		if err:= decoder.Decode(&v);err!=nil{
			log.Print(err)
			return

		}
		for k:= range v{
			if k=="Age"{
				delete(v,k)
			}
		}
		if err:= encoder.Encode(&v);err!=nil{
			log.Println(err)
		}
	}

}