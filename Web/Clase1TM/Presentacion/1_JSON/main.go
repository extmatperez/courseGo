package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	type product struct {
		Name      string
		Price     int
		Published bool
	}
	p := product{
		Name:      "MacBook Pro",
		Price:     1500,
		Published: true,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	p1 := product{}

	if err := json.Unmarshal([]byte(jsonData), &p1); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p1)

}
