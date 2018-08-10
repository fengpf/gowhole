package main

import (
	"fmt"
)

type Base struct{}
type Baser interface {
	Get() float32
}

type TypeOne struct {
	value float32
	Base
}

type TypeTwo struct {
	value float32
	Base
}

type TypeThree struct {
	value float32
	Base
}

func (t *TypeOne) Get() float32 {
	return t.value
}

func (t *TypeTwo) Get() float32 {
	return t.value * t.value
}

func (t *TypeThree) Get() float32 {
	return t.value + t.value
}

func main() {
	base := Base{}
	t1 := &TypeOne{1, base}
	t2 := &TypeTwo{2, base}
	t3 := &TypeThree{4, base}

	bases := []Baser{Baser(t1), Baser(t2), Baser(t3)}

	for s, _ := range bases {
		switch bases[s].(type) {
		case *TypeOne:
			fmt.Println("TypeOne")
		case *TypeTwo:
			fmt.Println("TypeTwo")
		case *TypeThree:
			fmt.Println("TypeThree")
		}

		fmt.Printf("The value is:  %f\n", bases[s].Get())
	}
}
