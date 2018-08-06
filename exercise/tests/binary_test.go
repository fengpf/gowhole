package tests

import (
	"fmt"
	"testing"
)

type Up struct {
	ID        int `json:"id"`
	Attribute int `json:"attribute"`
}

func (u *Up) AttrSet(v int, bit uint) {
	u.Attribute = u.Attribute&(^(1 << bit)) | (v << bit)
}

// AttrVal get attribute.
func (u *Up) AttrVal(bit uint) int {
	return (u.Attribute >> bit) & int(1)
}

func Test_set(t *testing.T) {
	from := uint(0)
	is := 0
	mdlUp := &Up{
		ID:        1,
		Attribute: 15,
	}
	// mdlUp = &Up{
	// 	ID: 1,
	// }
	mdlUp.AttrSet(is, from)
	fmt.Println(mdlUp)
}
