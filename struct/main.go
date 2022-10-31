package main

import "fmt"

func main() {
	/*var alex person
	alex.age = 20
	alex.firstName = "manish"
	alex.lastName = "paryani"*/

	//person1 := model.Person{"manish", "paryani", 25, model.ContactInfo{email: "manihs", zipCode: 43}}
	person1 := Person{"manish", "paryani", 10, ContactInfo{"manish", 234}}
	person1.print()
	pointerToperson1 := &person1
	pointerToperson1.updateName("ajay")
	person1.print()

	colors := make(map[string]string)
	/*colors := map[string]string{
		"red":  "#ff000",
		"blue": "#bb000",
	}*/
	colors["white"] = "0000"
	for s, v := range colors {
		fmt.Println(s)
		fmt.Println(v)
	}
	fmt.Println(colors)
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p *Person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
