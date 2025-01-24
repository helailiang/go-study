package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type Address struct {
	City    string
	Country string
}

type Contact struct {
	Email  string
	Phones []string
}

func (c *Contact) String() string {
	fmt.Printf("================================================================\n")
	return fmt.Sprintf("%#v", c)
}

type Employee struct {
	Name      string
	Age       int32
	Addresses []Address
	Contact   *Contact
}

type Manager struct {
	Name      string
	Age       int32
	Addresses []Address
	Contact   *Contact
}

func main() {
	employee := Employee{
		Name: "John Doe",
		Age:  30,
		Addresses: []Address{
			{City: "New York", Country: "USA"},
			{City: "San Francisco", Country: "USA"},
		},
		Contact: &Contact{
			Email:  "123@163.com",
			Phones: []string{"11111", "22222"},
		},
	}

	manager := Manager{}

	//copier.CopyWithOption(&manager, &employee, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	copier.Copy(&manager, &employee)
	fmt.Printf("Manager: %#v\n", manager)
	fmt.Printf("Manager.Contact: %#v\n", manager.Contact)
	// Output: Manager struct showcasing copied fields from Employee,
	// including overridden and deeply copied nested slices.
}
