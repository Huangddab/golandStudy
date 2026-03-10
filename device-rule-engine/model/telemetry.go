package model

/*
设备1

温度 65
湿度 80
气体 10
*/

type Telemetry struct {
	DeviceId    string
	Temperature float64
	Humidity    float64
	Gas         float64
}
