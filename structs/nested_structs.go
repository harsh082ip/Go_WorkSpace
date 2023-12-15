package main

import "fmt"

// example of nested structs
type person struct {
	name string
	age  int
	ad   address
}

type address struct {
	house_no   int
	firstLine  string
	secondLine string
}

func main() {
	p := person{}
	p.name = "Harsh"
	p.age = 18
	p.ad.house_no = 937
	p.ad.firstLine = "Somewhere In City"
	p.ad.secondLine = "Uttar Pradesh"

	fmt.Printf("Name: %v ,\n Age: %v ,\n ----Address----- \n House No: %v, \n First Line %v \n Second Line %v \n", p.name, p.age, p.ad.house_no, p.ad.firstLine, p.ad.secondLine)
}
