package main
import "fmt"

func main(){
	ids := []int{1,2,3,5,33,46,75,22,9,54,345}

	//loop through ids
	for i, id := range ids {
		fmt.Printf("%d -- ID %d\n", i, id)
	}

	//not using index
	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}

	//add IDS together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum: ", sum)

	//range with maps
	ages := map[string]int{"Bob":23, "Sherman":28, "Mike": 32}

	for key, value := range ages {
		fmt.Printf("%s: %d\n", key, value)
	}
}