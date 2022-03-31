package main

import "fmt"

func main() {
	dict := make(map[string]string)
	fmt.Println(dict)

	dict["ANK"] = "ANKARA"
	dict["iST"] = "İSTANBUL"
	dict["İZM"] = "İZMİR"
	fmt.Println(dict)
	ankara := dict["ANK"]
	fmt.Println("Seçilen şehir:", dict["ANK"], ankara)

	delete(dict, "İZM")
	fmt.Println(dict)

	dict["BRS"] = "BURSA"
	fmt.Println(dict)

	for k, v := range dict {
		fmt.Printf("%v : %v\n", k, v)
	}
	//dict map nesnesin elemanı adedince kapaasiteli key listesi oluşturma
	keys := make([]string, len(dict))
	i := 0
	for k := range dict {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	//key listesindeki index değerlerine göre dict nesnesinde bulunan şehirlerin yazımı

	for i := range keys {
		fmt.Println(dict[keys[i]])
	}

}
