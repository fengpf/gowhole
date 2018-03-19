package main

import "fmt"

//Rect for rect
type Rect struct {
	X, Y, Area, Length float64
}

//CallArea for call rect area
func (r *Rect) CallArea() {
	r.Area = r.X * r.Y
	fmt.Println(r)
}

//CallLength for call rect length
func (r Rect) CallLength() {
	r.Length = 2 * (r.X + r.Y)
	fmt.Println(r)
}

func main() {
	r1 := Rect{3, 4, 0, 0}
	r1.CallArea()
	fmt.Println(r1.Area)
	r1.CallLength()
	fmt.Println(r1.Length)
}
