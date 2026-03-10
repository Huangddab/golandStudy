package main

import (
	"device-rule-engine/engine"
	"device-rule-engine/model"
	"fmt"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

func main() {
	telemetry := &model.Telemetry{
		DeviceId:    "device001",
		Temperature: 65,
		Humidity:    50,
		Gas:         5,
	}

	alarm := &model.Alarm{}

	ruleFiles := []string{
		"rules/temperature.grl",
		"rules/gas.grl",
		"rules/humidity.grl",
	}

	ruleEngine := engine.NewRuleEngine(ruleFiles)

	dataContext := ast.NewDataContext()

	dataContext.Add("Telemetry", telemetry)
	dataContext.Add("Alarm", alarm)

	err := ruleEngine.Execute(dataContext)

	if err != nil {
		panic(err)
	}

	fmt.Println("设备:", telemetry.DeviceId)
	fmt.Println("报警类型:", alarm.Type)
	fmt.Println("报警等级:", alarm.Level)
	fmt.Println("报警信息:", alarm.Message)
}
