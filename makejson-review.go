package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

func main() {

	values := make(map[string]string)

	fmt.Printf("Enter a name: ")
	
	name := bufio.NewScanner(os.Stdin)

	name.Scan()

	values["name"] = name.Text()

	fmt.Printf("Enter an address: ")

	address := bufio.NewScanner(os.Stdin)

	address.Scan()

	values["address"] = address.Text()

	dat, e := json.Marshal(values)

	if e != nil {
		panic(e)
	}

	fmt.Println(string(dat))
	

	




}