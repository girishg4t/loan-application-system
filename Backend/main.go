package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/loan-application-system/pkg/account_software"
	"github.com/loan-application-system/pkg/decision_engine"
	"github.com/loan-application-system/pkg/handler"
	"github.com/loan-application-system/pkg/middleware"
	"github.com/loan-application-system/pkg/model"
	"github.com/loan-application-system/pkg/redis"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}
	rdb, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		log.Fatalln("failed to get redis config")
	}
	as := account_software.NewAccountSoftware(context.Background(), redis.Config{
		Database: rdb,
		Address:  os.Getenv("REDIS_ADDRESS"),
	})
	de := decision_engine.NewDecisionEngine()

	m := middleware.AuthProvider{Config: model.Config{
		ApiKey: os.Getenv("API_KEY"),
	}}
	r := mux.NewRouter()
	h := handler.NewUserHandler(as, de)
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(m.Authenticate)
	api.Path("/{businessName}/balancesheet/{accProvider}").HandlerFunc(h.HandleBalanceSheet).Methods(http.MethodPost)
	api.Path("/{businessName}/submit").HandlerFunc(h.HandleSubmitApplication).Methods(http.MethodPost)

	port := os.Getenv("PORT")
	log.Println("Running local on port: ", port)
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"X-API-KEY"},
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},

		AllowCredentials: true,
	})

	handler := c.Handler(api)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), handler))

}
