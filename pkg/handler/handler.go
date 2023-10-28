package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	accountsoftware "github.com/loan-application-system/pkg/account-software"
	decisionengine "github.com/loan-application-system/pkg/decision-engine"
	"github.com/loan-application-system/pkg/model"
	"github.com/loan-application-system/pkg/rules"
)

type UserHandler struct {
	re rules.RuleEngine
	as accountsoftware.IAccountSoftware
	de decisionengine.IDecisionEngine
}

func NewUserHandler(re rules.RuleEngine, as accountsoftware.IAccountSoftware, de decisionengine.IDecisionEngine) UserHandler {
	return UserHandler{
		re: re,
		as: as,
		de: de,
	}

}

func (h UserHandler) HandleBalanceSheet(w http.ResponseWriter, req *http.Request) {
	log.Println("handling user request for loan application")
	accProvider := mux.Vars(req)["accProvider"]
	if accProvider == "" {
		http.Error(w, "invalid account provider", http.StatusBadRequest)
		return
	}

	bs := h.as.GetBalanceSheet(accProvider)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(bs)
}

func (h UserHandler) HandleSubmitApplication(w http.ResponseWriter, req *http.Request) {
	log.Println("handling user request for loan application")

	var u model.UserApplication
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if u.LoanAmount <= 0 {
		err := fmt.Errorf("loan amount is not valid")
		_ = json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bs := h.as.GetBalanceSheet(u.AccountProvider)

	out := h.re.GetSummary(u, bs)
	decision := h.de.MakeDecision(out)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(model.FinalOutcome{
		Decision:       decision,
		ApprovedAmount: (out.PreAssessment / 100) * u.LoanAmount,
	})
}
