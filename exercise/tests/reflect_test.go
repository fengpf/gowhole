package tests

import (
	"fmt"
	"reflect"
	"testing"

	"math/rand"
)

func Test_val(t *testing.T) {
	i := 1
	ty := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	name := reflect.Indirect(v).Type().Name()
	fmt.Println(v, ty, name)

	setMethods(ty)
}

func setMethods(ty reflect.Type) {
	methods := map[string]reflect.Method{}
	fmt.Println(2222, ty.NumMethod())
	for m := 0; m < ty.NumMethod(); m++ {
		println(111, m)
		method := ty.Method(m)
		mtype := method.Type
		mname := method.Name
		fmt.Printf("%v pp %v ni %v 1k %v 2k %v no %v\n",
			mname, method.PkgPath, mtype.NumIn(), mtype.In(1).Kind(), mtype.In(2).Kind(), mtype.NumOut())
		if method.PkgPath != "" || // capitalized?
			mtype.NumIn() != 3 ||
			//mtype.In(1).Kind() != reflect.Ptr ||
			mtype.In(2).Kind() != reflect.Ptr ||
			mtype.NumOut() != 0 {
			// the method is not suitable for a handler
			//fmt.Printf("bad method: %v\n", mname)
		} else {
			// the method looks like a handler
			methods[mname] = method
		}
	}
	fmt.Println(methods, (rand.Int() % 1000), (rand.Int()%1000) < 100)
}

func Test_updateValue(t *testing.T) {
	x := 2                   // value type variable?
	a := reflect.ValueOf(2)  // 2 int no
	b := reflect.ValueOf(x)  // 2 int no
	c := reflect.ValueOf(&x) // &x *int no
	d := c.Elem()            // 2 int yes (x)

	fmt.Println(a.CanAddr()) // "false"
	fmt.Println(b.CanAddr()) // "false"
	fmt.Println(c.CanAddr()) // "false"
	fmt.Println(d.CanAddr()) // "true

	e := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := e.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3

	fmt.Println(a, b, c, d)

	e.Set(reflect.ValueOf(4))
	fmt.Println(x) // "4"

}
