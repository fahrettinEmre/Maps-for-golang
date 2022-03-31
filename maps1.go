package main

import "fmt"

func main() {
	/* a := []int{1, 2, 3}
	/* array = append(array, 10.45, 45.2)
	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		fmt.Printf("%d is %f\n", i, array[i])
	}
	*/
	/* a = append(a, 7)
	fmt.Println(a) */

	dict := make(map[string]int)
	dict["Ankara"] = 06
	dict["istanbul"] = 34
	dict["Yalova"] = 77

	//fmt.Println(dict["Ankara"])
	delete(dict, "Yalova")
	fmt.Println(dict)

}
