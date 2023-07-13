package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

func main() {
	db, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/dbname")
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}
	defer db.Close(context.Background())

	stripe.Key = "your-stripe-key"

	r := gin.Default()

	r.POST("/payment", func(c *gin.Context) {
		// Extract payment information from the request body

		// Make payment through Stripe
		_, err := charge.New(&stripe.ChargeParams{
			Amount:   stripe.Int64(2000),
			Currency: stripe.String(string(stripe.CurrencyUSD)),
			Source:   &stripe.SourceParams{Token: stripe.String("tok_visa")},
		})
		if err != nil {
			// Handle error
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment"})
			return
		}

		// Store payment information in the database

		c.JSON(http.StatusOK, gin.H{"status": "Payment successful"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
