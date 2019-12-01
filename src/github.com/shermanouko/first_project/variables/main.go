package main
import "fmt"

func main(){
	//declare variable
	var name = "Sherman"
	var age = 24
	const gender = "male"

	//method 2 of declaring variables
	second_name := "Ouko"
	height := 1.6
	course, year := "Computer Engineering", 5

	//print out
	fmt.Println(name, age, gender, second_name, height, course, year)

	//get type
	fmt.Printf("%T\n", name)
}