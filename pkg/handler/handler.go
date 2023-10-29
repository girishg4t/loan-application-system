package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/loan-application-system/pkg/account_software"
	"github.com/loan-application-system/pkg/decision_engine"
	"github.com/loan-application-system/pkg/model"
	"github.com/loan-application-system/pkg/rules"
)

type UserHandler struct {
	as account_software.IAccountSoftware
	de decision_engine.IDecisionEngine
}

func NewUserHandler(as account_software.IAccountSoftware, de decision_engine.IDecisionEngine) UserHandler {
	return UserHandler{
		as: as,
		de: de,
	}
}

func (h UserHandler) HandleBalanceSheet(w http.ResponseWriter, req *http.Request) {
	log.Println("handling user request for loan application")

	bn := mux.Vars(req)["businessName"]
	if bn == "" {
		http.Error(w, "business name is empty", http.StatusBadRequest)

		return
	}

	ap := mux.Vars(req)["accProvider"]
	if ap == "" {
		http.Error(w, "account provider is empty", http.StatusBadRequest)

		return
	}

	bs := h.as.GetBalanceSheet(req.Context(), model.UserApplication{AccountProvider: ap, BusinessName: bn})

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
	bs := h.as.GetBalanceSheet(req.Context(), u)

	rule := rules.NewRuleEngine(bs)
	out := rule.RequestOutcome(u)
	decision := h.de.MakeDecision(req.Context(), out)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(model.FinalOutcome{
		Decision:       decision,
		ApprovedAmount: (out.PreAssessment / 100) * u.LoanAmount,
	})
}
