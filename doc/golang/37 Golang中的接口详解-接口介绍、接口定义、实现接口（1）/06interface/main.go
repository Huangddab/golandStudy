package main

import "fmt"

// 接口
type Usber interface {
	start()
	stop()
}

//电脑 结构体
type Computer struct {
}

// 电脑的工作方法 接受一个usb接口类型的参数
func (c Computer) work(usb Usber) {
	//要判断usb的类型
	if _, ok := usb.(Phone); ok { // 判断usb的类型是否是Phone类型
		usb.start() //调用手机的方法
	} else {
		usb.stop() //调用相机的方法
	}

}

//手机
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

//照相机
type Camera struct {
}

func (p Camera) start() {
	fmt.Println("相机启动")
}
func (p Camera) stop() {
	fmt.Println("相机关机")
}

func main() {

	var computer = Computer{} //实例化一个电脑类型的变量
	var phone = Phone{        //实例化一个手机类型的变量
		Name: "小米手机",
	}
	var camera = Camera{} //实例化一个照相机类型的变量

	computer.work(phone)  //电脑工作  传入一个手机类型的变量  实现了Usber接口的方法  调用了手机的方法
	computer.work(camera) //传入一个照相机类型的变量  实现了Usber接口的方法  调用了照相机的方法

}
