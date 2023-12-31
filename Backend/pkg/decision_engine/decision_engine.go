package decision_engine

import (
	"context"

	"github.com/loan-application-system/pkg/model"
)

type DecisionEngine struct {
}

type IDecisionEngine interface {
	MakeDecision(ctx context.Context, r model.FinalOutput) bool
}

func NewDecisionEngine() DecisionEngine {
	return DecisionEngine{}
}

// MakeDecision will make the decision based on preAssessment value
func (d DecisionEngine) MakeDecision(ctx context.Context, r model.FinalOutput) bool {
	return r.PreAssessment > 20
}
