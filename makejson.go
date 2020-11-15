package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	var data = make(map[string]interface{})
	fmt.Print("Enter your name: ")
	if _, err := fmt.Scan(&name); err == nil {
		data["name"] = name
	}
	fmt.Print("Enter your address: ")
	if _, err := fmt.Scan(&address); err == nil {
		data["address"] = address
	}
	jsonObject, _ :=json.Marshal(data)
	fmt.Println(string(jsonObject))
}
