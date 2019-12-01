package main
import "fmt"

func main(){
	//Arrays
	var fruitArr [2]string

	//assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	//declare and assign
	carArr := [2]string{"BMW", "Merc"}

	fmt.Println(fruitArr)
	fmt.Println(carArr)

	//slices
	fruitSlice := []string{"APple", "Orange", "Banana"}

	fmt.Println(fruitSlice)

}