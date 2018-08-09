package slice

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"testing"
	"unicode/utf8"
)

func Test_main(t *testing.T) {
	a()
	println("----------------------")
	b()

	//i如果省略，默认是0；j如果省略默认是原数组或者切片的长度,所以例子中的三个新切片的值是一样的。这里注意的是i和j都不能超过原切片或者数组的索引。
	slice := []int{1, 2, 3, 4, 5}
	slice1 := slice[:]
	slice2 := slice[0:]
	slice3 := slice[:5]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	//新的切片和原切片共用的是一个底层数组，所以当修改的时候，底层数组的值就会被改变，所以原切片的值也改变了。当然对于基于数组的切片也一样的。
	newSlice := slice[1:3]
	newSlice[0] = 10
	fmt.Println(slice)
	fmt.Println(newSlice)
	//对于底层数组容量是k的切片slice[i:j]来说 长度：j-i 容量:k-i
	newSlice2 := slice[1:3]
	fmt.Printf("newSlice长度:%d,容量:%d\n\r", len(newSlice2), cap(newSlice2))
	//这样我们就创建了一个长度为2-1=1，容量为3-1=2的新切片,不过第三个索引，不能超过原切片的最大索引值5
	newSlice3 := slice[1:2:3]
	fmt.Printf("newSlice长度:%d,容量:%d\n\r", len(newSlice3), cap(newSlice3))
	//切片只能访问到其长度内的元素，访问超过长度外的元素，会导致运行时异常，与切片容量关联的元素只能用于切片增长
	fmt.Println(slice[2]) //获取值
	slice[2] = 10         //修改值
	fmt.Println(slice[2]) //输出10
	//切片算是一个动态数组，所以它可以按需增长，我们使用内置append函数即可。
	//append函数可以为一个切片追加一个元素，至于如何增加、返回的是原切片还是一个新切片、长度和容量如何改变这些细节，append函数都会帮我们自动处理。
	newSlice4 := slice[1:3]
	newSlice4 = append(newSlice4, 12)
	newSlice4 = append(newSlice4, 13)
	fmt.Printf("newSlice长度:%d,容量:%d\n\r", len(newSlice4), cap(newSlice4))
	newSlice4 = append(newSlice4, 14)
	fmt.Printf("newSlice长度:%d,容量:%d\n\r", len(newSlice4), cap(newSlice4))
	newSlice4 = append(newSlice4, 15)
	fmt.Println(newSlice4)
	fmt.Println(slice)
}

func a() {
	x := []int{}
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 5)

	fmt.Printf("x---ptr(%p),len(%d),cap(%d),val(%v)\n", x, len(x), cap(x), x)
	y := append(x, 3)
	fmt.Printf("y---ptr(%p),len(%d),cap(%d),val(%v)\n", y, len(y), cap(y), y)
	y = append(y, 8)

	z := append(x, 4)
	fmt.Printf("y---ptr(%p),len(%d),cap(%d),val(%v)\n", y, len(y), cap(y), y)
	fmt.Printf("z---ptr(%p),len(%d),cap(%d),val(%v)\n", z, len(z), cap(z), z)
}

func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	fmt.Printf("x---ptr(%p),len(%d),cap(%d),val(%v)\n", x, len(x), cap(x), x)
	y := append(x, 3)
	fmt.Printf("y---ptr(%p),len(%d),cap(%d),val(%v)\n", y, len(y), cap(y), y)
	z := append(x, 4)
	fmt.Printf("y---ptr(%p),len(%d),cap(%d),val(%v)\n", y, len(y), cap(y), y)
	fmt.Printf("z---ptr(%p),len(%d),cap(%d),val(%v)\n", z, len(z), cap(z), z)
}

func aa() {
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

func bb() {
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
// 	type Tags struct {
// 		string
// 	}
// 	type Labels struct {
// 		*Tags
// 	}
// 	type app struct {
// 		Labels *Labels
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

func Test_sliceDuration(t *testing.T) {
	var (
		ds []int64
		d  int64
		i  int64
	)
	d = int64(90)
	for i = 1; i < 6; i++ {
		j := float64(d / 5 * i)
		ds = append(ds, int64(math.Ceil(j)))
	}
	fmt.Println(ds)
}

func Test_sliceCunt(t *testing.T) {
	str := "HelloWord"
	l1 := len([]rune(str))
	l2 := bytes.Count([]byte(str), nil)
	l3 := strings.Count(str, "")
	l4 := utf8.RuneCountInString(str)
	fmt.Println("l1", l1)
	fmt.Println("l2", l2)
	fmt.Println("l3", l3)
	fmt.Println("l4", l4)

	// func Count(s, sep string) int{}
	// 判断字符sep在字符串s中出现的次数，没有找到则返回-1，如果为空字符串("")则返回字符串的长度+1
	fmt.Println(strings.Count(str, "o"))

	str1 := "Hello, 世界"
	fmt.Println(len(str))                     //打印结果：9  （如果是纯英文字符的字符串，可以使用来判断字符串的长度）
	fmt.Println(len(str1), len([]rune(str1))) // 打印结果：13
}
