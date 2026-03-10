package main

import (
	"fmt"
	"time"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type MyFact struct {
	ID               string
	IntAttribute     int64
	StringAttribute  string
	BooleanAttribute bool
	FloatAttribute   float64
	TimeAttribute    time.Time
	WhatToSay        string
}

func (mf *MyFact) GetWhatToSay(sentence string) string {
	return fmt.Sprintf("Let say \"%s\"", sentence)
}

func main() {

	// 创建事实
	myFact := &MyFact{
		ID:               "chemr",
		IntAttribute:     123,
		StringAttribute:  "Some string value",
		BooleanAttribute: true,
		FloatAttribute:   1.234,
		TimeAttribute:    time.Now(),
	}

	// DB:= model.Alarm{}

	// func (db*DB) loadValue() float64{

	// 	return 50.3
	// }

	// dataCtx.Add("DB", DB)

	// 加载事实到上下文
	dataCtx := ast.NewDataContext()

	// err := dataCtx.AddJSON("MF", "JSON")

	err := dataCtx.Add("MF", myFact)
	if err != nil {
		panic(err)
	}

	// 创建知识库
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	// 加载规则
	fileRes := pkg.NewFileResource("rules.grl")
	err = ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", fileRes)
	if err != nil {
		panic(err)
	}

	// 获取知识库实例
	knowledgeBase, err := knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")
	if err != nil {
		panic(err)
	}

	// 执行引擎
	engine := engine.NewGruleEngine()

	err = engine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		panic(err)
	}

	// 打印结果
	fmt.Println(myFact.WhatToSay)
	fmt.Println(myFact.IntAttribute)
}
