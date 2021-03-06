package base

import (
	"gengine/context"
	"gengine/core/errors"
)

type RuleContent struct {
	Statements        *Statements
	knowledgeContext  *KnowledgeContext
	dataCtx           *context.DataContext
}

func (t *RuleContent) Initialize(kc *KnowledgeContext, dc *context.DataContext) {
	t.knowledgeContext = kc
	t.dataCtx = dc
	if t.Statements!=nil {
		t.Statements.Initialize(kc, dc)
	}
}

func (t *RuleContent) Execute(Vars map[string]interface{}) error {
	_, err := t.Statements.Evaluate(Vars)
	if err != nil {
		return err
	}
	return nil
}

func (t *RuleContent)AcceptStatements(stmts *Statements)error {
	if t.Statements == nil{
		t.Statements = stmts
		return nil
	}
	return errors.Errorf("RuleContent's statements set twice.")
}


