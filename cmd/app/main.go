package main

import (
	"fmt"
	"log"
	"os"
	"payment_service/internal/payment"
	"payment_service/internal/pkg/common"
	"payment_service/internal/pkg/postgres"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load("deploy/.env")
	if err != nil {
		log.Fatalf("cannot open .env file: %sv", err)
	}

	postgresDB, cancelPostgres, err := postgres.OpenPostgresDB(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB")))
	if err != nil {
		log.Fatal("cannot open postgres connection ", err)
	}
	defer cancelPostgres(postgresDB)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(common.NewErrorHandler())

	postgresRepo := payment.NewPostgresRepository(postgresDB)
	service := payment.NewService(postgresRepo)
	handler := payment.NewHandler(service)

	apiRouter := router.Group("/api")
	{
		apiRouter.POST("/send", handler.MakeTransaction)
		apiRouter.GET("/transactions", handler.GetLastTransactions)
		apiRouter.GET("/wallet/:address/balance", handler.GetBalance)
	}

	addr := ":" + os.Getenv("SERVICE_PORT")
	err = router.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
