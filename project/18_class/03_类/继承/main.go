package main

import "fmt"

type Human struct {
	name   string
	age    int
	gender string
}

// 定义一个方法
func (p *Human) Eat() {
	fmt.Println("this is:" + p.name)
}

// 定义一个学生类 嵌套一个Human类
type Student1 struct {
	hum    Human // 包含Human类型的变量
	school string
}

// 定义一个老师 继承Human
type Teacher struct {
	Human          // 直接继承Human类 没有字段名称
	subject string //学科
}

func main() {
	s1 := Student1{
		hum: Human{
			name:   "lily",
			age:    18,
			gender: "女",
		},
		school: "清华大学",
	}
	fmt.Println("学生的姓名是:", s1.hum.name)
	// 修改
	s1.hum.name = "jony"
	fmt.Println("学生的姓名是:", s1.hum.name)

	t1 := Teacher{}
	fmt.Println("老师的姓名是:", t1.name)
	t1.name = "Tom"
	fmt.Println("老师的姓名是:", t1.name)

}
