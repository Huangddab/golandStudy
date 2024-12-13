package main

import "fmt"

// golang中空接口和类型断言的使用细节

// 定义一个结构体
type Address struct {
	City, Province string
	number         int
}

func main() {
	var userinfo = make(map[string]interface{})
	userinfo["name"] = "张三"
	userinfo["age"] = 18
	userinfo["hobby"] = []string{"篮球", "足球"} // 切片

	fmt.Println(userinfo)          // map[age:18 hobby:[篮球 足球] name:张三]
	fmt.Println(userinfo["hobby"]) // [篮球 足球]
	// fmt.Println(userinfo["hobby"][1]) // (map index expression of type interface{})
	fmt.Println(userinfo["hobby"].([]string)[0]) // 篮球
	fmt.Println(userinfo["name"])                // 张三

	var address = Address{"武汉", "湖北", 123456} // 实例化一个结构体

	userinfo["address"] = address // 将结构体赋值给map的value

	fmt.Println(userinfo["address"]) //	{武汉 湖北 123456}

	hobby2, _ := userinfo["hobby"].([]string) // 类型断言
	// 将interface{}类型的value转换为[]string类型的value   hobby是一个切片类型的变量
	hobby2 = append(hobby2, "乒乓球") // 追加一个元素
	fmt.Println(hobby2)            // [篮球 足球 乒乓球]
	fmt.Println(userinfo["hobby"]) // [篮球 足球]

	address2, _ := userinfo["address"].(Address) // 类型断言
	// 将interface{}类型的value转换为Address类型的value
	fmt.Println(address2.City, address2.Province, address2.number)
}
