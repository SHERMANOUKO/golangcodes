package main
import "fmt"

//a map is a ley value pair
func main(){
	emails := make(map[string]string)

	//assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sherman"] = "sherman@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	
	//delete from map
	delete(emails, "Mike")
	fmt.Println(emails)

	//declare maps and add key values
	ages := map[string]int{"Bob":23, "Sherman":28, "Mike": 32}
	fmt.Println(ages)
}