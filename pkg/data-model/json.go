package datamodel

import (
	"encoding/json"
	"fmt"
)

// User is the data model for user entity
type User struct {
	Name    string  `json:"name"`
	Age     uint8   `json:"age"`
	Address Address `json:"address"`
}

// Address is the data mdoel for any address that can be associated to entities
type Address struct {
	Street string `json:"street"`
	Zip    string `json:"zip"`
}

// PlayWithJSON like a king
func PlayWithJSON() {

	address := Address{Street: "Alexanderstraße 5", Zip: "10178"}

	john := User{Name: "John", Age: 25, Address: address}

	johnToJSON, _ := json.Marshal(john)

	fmt.Println()
	fmt.Printf("Stringified John ---> %s", string(johnToJSON))
	fmt.Println()
	fmt.Println()

	var johnFromJSON User

	json.Unmarshal(johnToJSON, &johnFromJSON)

	fmt.Println()
	fmt.Printf("Parsed John ---> %+v", johnFromJSON)
	fmt.Println()
	fmt.Println()
}
