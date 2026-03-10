package model

/*
设备1
温度报警
HIGH
温度超过60℃
*/

type Alarm struct {
	DeviceId string
	Type     string
	Level    string
	Message  string
}
