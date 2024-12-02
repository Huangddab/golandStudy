// 2、请求出一个数组的最大值，并得到对应的下标
// 1、声明一个数组 var intArr[5] = [...]int {1, -1, 12, 65, 11}
// 2、假定第一个元素就是最大值，下标就 0
// 3、然后从第二个元素开始循环比较，如果发现有更大，则交换
// package main

// import "fmt"

// func main() {
// 	var intArr = [...]int{1, -1, 112, 65, 11}
// 	max := intArr[0]
// 	for i := 0; i < len(intArr); i++ {
// 		fmt.Printf("item=%v\t", intArr[i])

// 		if max < intArr[i] {
// 			max = intArr[i]
// 		}
// 		fmt.Printf("maxItem=%v\n", max)

// 	}
// 	fmt.Printf("max=%v\t", max)
// }

// 从数组[1, 3, 5, 7, 8]中找出和为 8 的两个元素的下标分别为(0,3)和(1,2)

package main

import "fmt"

func main() {
	var intArr = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(intArr)-1; i++ {
		for j := i + 1; j < len(intArr); j++ {
			if intArr[i]+intArr[j] == 8 {
				fmt.Printf("(%v,%v)\n", i, j)
			}
		}
	}
}
