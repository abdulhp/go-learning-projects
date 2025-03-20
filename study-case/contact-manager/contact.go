package main

import "fmt"

type Contact struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Address string `json:"address"`
}

var contacts []Contact

func list() {
	loadContacts()
	if len(contacts) == 0 {
		fmt.Println("No contacts saved")
	}

	fmt.Println("Total Contacts: ", len(contacts))
	fmt.Println("-----------------------")

	for _, contact := range contacts {
		fmt.Println("ID: ", contact.ID)
		fmt.Println("Name: ", contact.Name)
		fmt.Println("Phone: ", contact.Phone)
		fmt.Println("Email: ", contact.Email)
		fmt.Println("Address: ", contact.Address)
		fmt.Println("-----------------------")
	}
}
