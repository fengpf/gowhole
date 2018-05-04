package basictests

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_val(t *testing.T) {
	i := 1
	v := reflect.ValueOf(i)
	ty := reflect.TypeOf(i)
	fmt.Println(v, ty)
}
