package main

import "fmt"

func main() {
	/* arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	} */

	/* arr := make(map[string](int))

	arr["a"] = 0
	arr["b"] = 1
	arr["c"] = 2

	for key, value := range arr {
		fmt.Println("key:", key, "value:", value)
	} */
	//fmt.Println("sonuc:", sum(2, 3))
	sonuc := sum(2, 4)
	fmt.Println("sonuc", sonuc)
}

func sum(x int, y int) int {
	return x + y
}
