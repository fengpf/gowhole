package aaa

type Human struct {
	Name   string
	Age    int
	Weight int
}

type Student struct {
	Human      //匿名字段
	Speciality string
}
