package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	accountsoftware "github.com/loan-application-system/pkg/account-software"
	decisionengine "github.com/loan-application-system/pkg/decision-engine"
	"github.com/loan-application-system/pkg/handler"
	"github.com/loan-application-system/pkg/middleware"
	"github.com/loan-application-system/pkg/model"
	"github.com/loan-application-system/pkg/rules"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	re := rules.NewRuleEngine()
	as := accountsoftware.NewAccountSoftware()
	de := decisionengine.NewDecisionEngine()

	m := middleware.AuthProvider{Config: model.Config{
		ApiKey: os.Getenv("API_KEY"),
	}}
	r := mux.NewRouter()
	h := handler.NewUserHandler(re, as, de)
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(m.Authenticate)
	api.Path("/balancesheet/{accProvider}").HandlerFunc(h.HandleBalanceSheet).Methods(http.MethodPost)
	api.Path("/submit").HandlerFunc(h.HandleSubmitApplication).Methods(http.MethodPost)

	port := os.Getenv("PORT")
	log.Println("Running local on port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
