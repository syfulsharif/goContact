package main

import (
	"fmt"
	"github.com/syfulsharif/goContact/contact"
)

func main() {
	// Map Struct
	cm := contact.NewManager()
	cm.AddContact(contact.Contact{
		Name:  "John Doe",
		Phone: "+88018868288",
		Email: "jd@email.com",
	})

	cm.AddContact(contact.Contact{
		Name:  "Fun Doe",
		Phone: "+88018868388",
		Email: "fd@email.com",
	})

	cm.AddContact(contact.Contact{
		Name:  "Ron Doe",
		Phone: "+88018868488",
		Email: "rd@email.com",
	})

	for index, c := range cm.ContactList() {
		fmt.Printf("Contact %d: %v\n", index, c)
	}
}
