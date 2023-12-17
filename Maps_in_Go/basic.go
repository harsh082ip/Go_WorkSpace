package main

import "fmt"

func main() {

	// Maps = map[key]value
	// it will not be ordered like Arrays
	// it is faster than Arrays

	var mp map[string]int = map[string]int{
		"id":       43242344,
		"password": 34231,
	}

	fmt.Println(mp)
	fmt.Println(mp["id"])

	// create a key
	mp["new_id"] = 400
	fmt.Println(mp)

	// update a key (same syntax)
	mp["id"] = 1234554
	fmt.Println(mp)

	// delete a key
	delete(mp, "new_id")
	fmt.Println(mp)

	// using make()
	m := make(map[string]int)
	m["hey"] = 1223

	// it will now override above value
	m = map[string]int{
		"id": 234323,
	}
	fmt.Println(m)
}
