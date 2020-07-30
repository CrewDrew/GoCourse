package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type PhoneBook struct {
	Name  string
	Phone int64
}

func main() {

	var phones []PhoneBook

	phones = append(phones, PhoneBook{Name: "Vasya", Phone: 79229552145})
	phones = append(phones, PhoneBook{Name: "Petya", Phone: 79214562745})
	phones = append(phones, PhoneBook{Name: "Kolya", Phone: 79222352545})

	jsonPhoneBook, err := json.Marshal(phones)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(jsonPhoneBook))

	ioutil.WriteFile("phonebook.json", jsonPhoneBook, 0644)

}
