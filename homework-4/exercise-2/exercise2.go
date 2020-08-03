package main

import (
	"fmt"
	"sort"
)

type phones []string

type contact struct {
	name       string
	phonesList phones
}

type contacts []contact

func (p phones) ViewList() {
	for i, phone := range p {
		fmt.Printf("\t %v) %s \n", i+1, phone)
	}
}

func (c contacts) Len() int {
	return len(c)
}

func (c contacts) Less(i, j int) bool {
	if c[i].name < c[j].name {
		return true
	}
	return false
}

func (c contacts) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c contacts) PrintAll() {
	for i, person := range c {
		fmt.Printf("Контакт №%v: %s\n", i+1, person.name)
		person.phonesList.ViewList()
	}
}

func main() {
	var adressBook contacts

	adressBook = append(adressBook, contact{name: "Vasya", phonesList: phones{"+79231551234", "+79231551235"}})
	adressBook = append(adressBook, contact{name: "Kolya", phonesList: phones{"+79231553423", "+79231554435"}})
	adressBook = append(adressBook, contact{name: "Petya", phonesList: phones{"+79231534234", "+79231551555"}})
	adressBook = append(adressBook, contact{name: "Angie", phonesList: phones{"+79231534234", "+79231551555"}})

	adressBook.PrintAll()

	sort.Sort(adressBook)

	adressBook.PrintAll()

}
