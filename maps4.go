package main

import "fmt"

func main() {
	//var sena int = 5
	//var kitap string =""
	//var karakter rune = 't'

	var listem []string
	// listem = append(listem, 1)
	// listem = append(listem, 2)
	// listem = append(listem, 3)
	// listem = append(listem, 4)
	// listem = append(listem, 5)
	// listem = append(listem, 6)
	listem = append(listem, "akif", "süha", "11", "2323", "deneme", "faho", "samet", "sena") //15 ,27,1003
	// v := make([]rune, 0)
	// v = append(v, 't', 'a', 't', 'a', 'r')
	// i := len(listem)
	// fmt.Println(listem[i-1])

	// Key-Value
	// 0 - 0. elemanconst
	// len(t)-1 sonuncu elemanc
	// len(t)/ ortadaki eleman

	// i := len(listem)
	// t := cap(listem)

	// fmt.Println(listem)
	// fmt.Println(i)
	// fmt.Println(t)
	// fmt.Println(v)

	// faoDickt := make(map[string]int, 0)
	// faoDickt["akif"] = 29
	// faoDickt["sena"] = 24
	// faoDickt["süha"] = 56

	stringMap := make(map[string]string, 0)
	stringMap["istanbul"] = "kartal,pendik,maltepe,tuzla"
	stringMap["paris"] = "massy,evry,cergy"

	fmt.Println(stringMap["istanbul"])

}
