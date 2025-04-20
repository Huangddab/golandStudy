package main

import "fmt"

// Person类 绑定方法 Eat Run Laugh 成员

type Person struct {
	// 成员属性
	name   string
	age    int
	gender string
	score  float64
}

// 方法 在类的外面绑定方法
func (p Person) Eat() {
	println(p.name + " is eating")
	p.name = "ddd" // 非指针无法修改
}

func (this *Person) Run() {
	println(this.name + " is running")
	this.name = "jony" // 指针的就可以修改
}

func main() {
	lily := Person{
		name:   "Lily",
		age:    59,
		gender: "女生",
		score:  99.9,
	}

	fmt.Println("Lily:", lily)
	lily.Eat()
	fmt.Println("没有使用 p *Person")
	lily.Run()
	fmt.Println("使用 p *Person")
	fmt.Println("Lily:", lily)

}
