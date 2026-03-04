package basic

import (
	"fmt"
	"math"
	"testing"
)

// 定义接口
type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 实现接口方法 - 注意返回类型必须是float64
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 接口的使用方式：
func measure(g geometry) {
	fmt.Printf("面积: %.2f\n", g.area())
	fmt.Printf("周长: %.2f\n", g.perim())
}

func TestInterface(t *testing.T) {
	r := rectangle{width: 3, height: 5}
	c := circle{radius: 2}

	// 方式1：直接调用结构体的方法
	fmt.Printf("长方形: 宽%.0f 高%.0f 面积%.0f 周长%.0f\n",
		r.width, r.height, r.area(), r.perim())
	fmt.Printf("圆形: 半径%.0f 面积%.2f 周长%.2f\n",
		c.radius, c.area(), c.perim())

	fmt.Println("\n--- 通过接口调用 ---")
	// 方式2：通过接口变量调用
	var g geometry

	// 长方形赋值给接口
	g = r
	fmt.Println("长方形（通过接口）：")
	measure(g)

	// 圆形赋值给接口
	g = c
	fmt.Println("圆形（通过接口）：")
	measure(g)

	// 方式3：接口切片
	fmt.Println("\n--- 接口切片示例 ---")
	shapes := []geometry{r, c}
	for i, shape := range shapes {
		fmt.Printf("形状 %d:\n", i+1)
		measure(shape)
		fmt.Println()
	}
}
