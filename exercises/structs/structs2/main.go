// structs2
// Make me compile!

package main

import "fmt"

type ContactDetails struct {
	phone string
}

type Person struct {
	// don't just create the phone field here. embed a new struct
	name string
	age int
	phone ContactDetails
}


func main() {
	contactDetails := ContactDetails{"4002 8922"}
	person := Person{name: "John", age: 32, phone: contactDetails}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
