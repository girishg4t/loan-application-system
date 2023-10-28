package decisionengine

import (
	"github.com/loan-application-system/pkg/model"
)

type DecisionEngine struct {
}

type IDecisionEngine interface {
	MakeDecision(r model.FinalOutput) bool
}

func NewDecisionEngine() DecisionEngine {
	return DecisionEngine{}
}

func (d DecisionEngine) MakeDecision(r model.FinalOutput) bool {
	return r.PreAssessment > 20
}
