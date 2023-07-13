package main

import (
	"context"
	"log"
	"net/http"

	service "github.com/brunomguimaraes/scrooge/pkg/payment"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type PaymentInfo struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Token    string `json:"token"`
}

func main() {
	db, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/dbname")
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}
	defer db.Close(context.Background())

	dbService := service.NewDBService(db)
	stripeService := service.NewStripeService("your-stripe-key")

	r := gin.Default()

	r.POST("/payment", func(c *gin.Context) {
		// Extract payment information from the request body
		var paymentInfo PaymentInfo
		if err := c.ShouldBindJSON(&paymentInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Make payment through Stripe
		_, err := stripeService.Charge(paymentInfo.Amount, paymentInfo.Currency, paymentInfo.Token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment"})
			return
		}

		// Store payment information in the database
		err = dbService.SavePaymentInfo(paymentInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save payment info"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Payment successful"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
