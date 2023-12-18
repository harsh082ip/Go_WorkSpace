package main

import "fmt"

type point struct {
	x int
	y int
}

func changeX(pt *point) {
	pt.x = 12
}

type circle struct {
	radius int
	p      *point
}

type withoutAVariable struct {
	r int
	// there is no need of an extra variable now to access values of point struct
	*point
}

func main() {
	p1 := &point{2, 3}
	fmt.Println(p1)
	fmt.Println(*p1)
	changeX(p1)
	fmt.Println(*p1)

	c := circle{2, &point{3, 5}}
	fmt.Println(c.p.x)
	// change x
	fmt.Println(c.p.x, c.p.y)
	c.p.x = 17
	fmt.Println(c.p.x, c.p.y)

	n := withoutAVariable{3, &point{3, 5}}
	fmt.Println(n.point)
	// change x of point
	fmt.Println(n.x)
	n.x = 87
	fmt.Println(n.x)

}
