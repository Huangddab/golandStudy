package main

import "fmt"

func main() {
	// /要求: 从控制台接收用户输入的信息,[姓名,年龄,薪水,是否通过考试]
	//方式1:使用fmt.Scanln()
	//方式2:使用fmt.Scanf()
	//方式1:使用fmt.Scanln()
	//1.先声明需要的变量
	var name string
	var age byte
	var sla float32
	var isPass bool
	//当程序执行到fmt.Scanln()时,程序会停止在这里,等待用户输入,并回车
	// fmt.Println("请输入姓名:")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水:")
	// fmt.Scanln(&sla)
	// fmt.Println("请输入是否通过考试:")
	// fmt.Scanln(&isPass)
	// fmt.Printf("姓名: %v, 年龄: %v, 薪水: %v, 是否通过考试: %v", name, age, sla, isPass)

	//方式2:fmt.Scanf(),可以按指定的格式输入
	fmt.Println("请输入姓名, 年龄, 薪水, 是否通过考试, 以空格隔开:")
	fmt.Scanf("%s %d %f %t", &name, &age, &sla, &isPass)
	fmt.Printf("姓名: %v, 年龄: %v, 薪水: %v, 是否通过考试: %v", name, age, sla, isPass)
}
