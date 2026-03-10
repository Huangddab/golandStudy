package main

import (
	"fmt"
	"log"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// Grule的核心流程：定义数据 -> 编写规则 -> 加载规则 -> 准备数据 -> 执行引擎
// 数据(事实 Fact) + 规则(Rule)  -> 规则引擎执行 -> 结果
/*
| 组件            | 作用              |
| ------------- | --------------- |
| Fact          | 业务数据（Go struct） |
| Rule          | 业务规则（类似Drools）  |
| KnowledgeBase | 规则集合            |
| DataContext   | 运行时数据           |
| RuleEngine    | 执行规则            |
*/
// 1.定义事实
type Purchase struct {
	UserLevel string  // 用户等级
	Amount    float64 // 数量
	Discount  float64 // 将计算出的折扣
}

func main() {
	// 2.创建事实对象
	myPurchase := &Purchase{
		UserLevel: "VIP",  
		Amount:    1000.0,
	}

	// 3.初始化知识库和规则构建器
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	// 4.定义您的业务规则
	// 5.构建编译规则

	resource := pkg.NewFileResource("rules/purchase.grl")
	err := ruleBuilder.BuildRuleFromResource("PurchaseRules", "1.0.0", resource)
	if err != nil {
		log.Fatal("规则编译失败：", err)
	}

	// 6.获取一个知识库实例
	knowledgeBase, err := knowledgeLibrary.NewKnowledgeBaseInstance("PurchaseRules", "1.0.0")
	if err != nil {
		log.Fatal("获取知识库实例失败：", err)
	}

	// 7.创建数据上下文并添加事实
	dataContext := ast.NewDataContext()
	// 将业务对象添加到上下文中，并赋予一个标识符
	err = dataContext.Add("Purchase", myPurchase) // 注意：这里的标识符"Purchase"对应规则中引用的名称
	if err != nil {
		log.Fatal("添加事实到上下文失败：", err)
	}

	// 8.创建规则引擎并执行
	ruleEngine := engine.NewGruleEngine()
	err = ruleEngine.Execute(dataContext, knowledgeBase)
	if err != nil {
		log.Fatal("规则执行失败：", err)
	}

	// 9.查看结果
	fmt.Printf("最终订单金额: %.2f\n", myPurchase.Amount)
	fmt.Printf("计算得出的折扣: %.2f\n", myPurchase.Discount)
	fmt.Printf("应付金额: %.2f\n", myPurchase.Amount-myPurchase.Discount)
}
