package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Coordinate() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
	var a Point = Point{2, 3}
	fmt.Println(a.Coordinate())
}
