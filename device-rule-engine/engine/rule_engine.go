package engine

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
)

type RuleEngine struct {
	knowledgeBase *ast.KnowledgeBase
	engine        *engine.GruleEngine
}

func NewRuleEngine(ruleFile []string) *RuleEngine {
	const (
		ruleSetName    = "DeviceRules"
		ruleSetVersion = "1.0.0"
	)

	// 创建知识库
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	// 加载多个规则
	for _, path := range ruleFile {
		// 注意：这里的入参是“规则文件路径”，必须用 NewFileResource 按文件读取
		resource := pkg.NewFileResource(path)
		err := ruleBuilder.BuildRuleFromResource(ruleSetName, ruleSetVersion, resource)
		if err != nil {
			logrus.Fatalf("加载规则失败 %s : %v", path, err)
		}
	}
	// 获取知识库实例
	kb, err := lib.NewKnowledgeBaseInstance(ruleSetName, ruleSetVersion)
	if err != nil {
		logrus.Fatalf("NewKnowledgeBaseInstance err: %v", err)
	}

	return &RuleEngine{
		knowledgeBase: kb,
		engine:        engine.NewGruleEngine(),
	}
}

func (r *RuleEngine) Execute(dataContext ast.IDataContext) error {
	return r.engine.Execute(dataContext, r.knowledgeBase)
}
