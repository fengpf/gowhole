package basictests

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic("wrong score")
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func Test_grade(t *testing.T) {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		// grade(101),
	)
}
