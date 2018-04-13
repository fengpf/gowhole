package base

import (
	"fmt"
	"testing"
)

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1) // commonTags := labelsToTags(app.Labels)
	fmt.Printf("%p\n", x)
	y := append(x, 2) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Printf("%p\n", y)
	z := append(x, 3) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Printf("%p\n", z)
	fmt.Println(y, z)
}

func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2) // commonTags := labelsToTags(app.Labels)
	fmt.Printf("%p\n", x)
	y := append(x, 3) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Printf("%p\n", y)
	z := append(x, 4) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Printf("%p\n", z)
	fmt.Println(y, z)
}

func TestFu(t *testing.T) {
	a()
	b()
}

// func testTag() {
// 	type RegistrationIntent struct {
// 		Name string
// 		Port int
// 		Tags []string
// 	}
// 	type app struct {
// 		Labels *[]Labels
// 	}
// 	type Labels struct {
// 		*Tags
// 	}
// 	commonTags := labelsToTags(app.Labels)
// 	var intents []RegistrationIntent
// 	for _, d := range definitions {
// 		intents = append(intents, RegistrationIntent{
// 			Name: app.labelsToName(d.Labels, nameSeparator),
// 			Port: task.Ports[d.Index],
// 			Tags: append(commonTags, labelsToTags(d.Labels)...), // ◀ Wrong tags here
// 		})
// 	}
// }

func labelsToTags(labels map[string]string) []string {
	tags := []string{}
	for key, value := range labels {
		if value == "tag" {
			tags = append(tags, key)
		}
	}
	return tags
}
