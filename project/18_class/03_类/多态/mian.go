package main

import "fmt"

// 实现go多态需要实现定义接口
//  人类的武器发起攻击 不同等级子弹效果不同

// 多态：同一个接口变量可以指向不同的具体类型实例
type IAttack interface {
	// 接口函数可以有多个 但是只能有函数原型 不可以有实现
	Attack()
}

type LowLevel struct {
	name  string
	level int
}
type HightLevel struct {
	name  string
	level int
}

func (a *LowLevel) Attack() {
	fmt.Println("我是:", a.name, "等级为:", a.level, "伤害值:", 100)
}

func (a *HightLevel) Attack() {
	fmt.Println("我是:", a.name, "等级为:", a.level, "伤害值:", 500)
}

// 定义一个多态的通用接口 传入不同的对象调用同样的接口
func DoAttack(a IAttack) {
	a.Attack()
}

func main() {
	// var player interface {} // 不能使用空接口
	// var player IAttack //定义一个包含Attract的接口变量

	LowLevel := LowLevel{
		name:  "小子弹",
		level: 1,
	}

	HightLevel := HightLevel{
		name:  "大子弹",
		level: 2,
	}
	LowLevel.Attack()
	HightLevel.Attack()

	// 对player赋值为lowlevel 接口需要用指针类型来赋值
	// player = &LowLevel
	// player.Attack() // 调用接口的函数

	// player = &HightLevel
	// player.Attack()

	DoAttack(&LowLevel)   // 传入接口
	DoAttack(&HightLevel) // 传入接口
}
