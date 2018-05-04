package basictests

import (
	"fmt"
	"testing"
)

type Rect struct {
	X, Y, Area, Length float64
}

func (r *Rect) CalArea() {

	r.Area = r.X * r.Y

}

func (r Rect) CalLength() {

	r.Length = 2 * (r.X + r.Y)

}

func Test_rect(t *testing.T) {

	r1 := Rect{3, 4, 0, 0}

	r1.CalArea()

	r1.CalLength()

	fmt.Println(r1.Area)

	fmt.Println(r1.Length)

	r2 := &Rect{3, 4, 0, 0}

	r2.CalArea()

	r2.CalLength()

	fmt.Println(r2.Area)

	fmt.Println(r2.Length)

}
