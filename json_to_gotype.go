package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person1 struct{
	Name string
	Age float64
	Location string
}
func main(){

j :=[]byte(`{"name":"Ramony","age":44,"location":"Texas"}`)

var p Person1

err := json.Unmarshal(j,&p)

if err != nil{
	log.Fatalf("Unabele to Decode The Json")

	
}
fmt.Printf("Name:%s\nAge:%d\nLocation:%s\n",p.Name,p.Age,p.Location)

}
 