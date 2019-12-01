package main
import ("fmt"
		"strconv")

//define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

//greeting method (value reciver)
func (p Person) greet() string {
	return "Hello my name is " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

//has bithday method (pointer recevier)
func (p *Person) hasBirthday() {
	p.lastName = "Ouko" 
}

//getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main(){
	//initialize person with struct
	person_1 :=  Person{
							firstName : "Sherman", 
							lastName : "Allen", 
							city : "Nairobi", 
							gender : "m",
							age : 24,
						}
	
	//alternative
	person_2 :=  Person{"Sherman 2", "Allen 2", "Nairobi", "f", 24}

	fmt.Println(person_1)
	fmt.Println(person_2)

	//get single field
	fmt.Println(person_1.firstName)

	//reaassign
	person_1.age++
	fmt.Println(person_1)

	fmt.Println(person_1.greet())

	person_1.hasBirthday()

	fmt.Println(person_1.greet())

	person_2.getMarried("Williams")
	person_1.getMarried("Williams")
	
	fmt.Println(person_1.greet())
	fmt.Println(person_2.greet())
}