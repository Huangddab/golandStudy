package main

import (
	"fmt"
	"strconv"
)

// string类型转基本数据类型

func main() {
	var str1 string = "true"
	var b1 bool

	// func strconv.ParseBool(str string) (bool, error)
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("b1 type is %T, b1=%v\n", b1, b1)
	// b1 type is bool, b1=true

	var str2 string = "1234546"
	var num1 int64
	// ParseInt(s string, base int, bitSize int) (i int64, err error)
	/*
	*返回字符串表示的整数值，接受正负号。
	*base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
	*bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	*返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRang
	 */
	num1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("num1 type is %T, num1=%v\n", num1, num1)
	// num1 type is int64, num1=1234546

	// 因为返回的是int64，如果需要int32类型，需要进行强制转换
	var num2 int
	num2 = int(num1)
	fmt.Printf("num2 type is %T, num2=%v\n", num2, num2)
	// num2 type is int, num2=1234546

	var str3 string = "12.67489"
	var fl64 float64
	// ParseFloat(s string, bitSize int) (float64, error)
	fl64, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("fl64 type is %T, fl64=%v\n", fl64, fl64)
	// fl64 type is float64, fl64=12.67489
}
