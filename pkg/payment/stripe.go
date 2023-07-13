package service

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

type StripeService struct {
	StripeKey string
}

func NewStripeService(stripeKey string) *StripeService {
	return &StripeService{
		StripeKey: stripeKey,
	}
}

func (s *StripeService) Charge(amount int64, currency, source string) (*stripe.Charge, error) {
	stripe.Key = s.StripeKey

	chargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		Source:   &stripe.SourceParams{Token: stripe.String(source)},
	}
	return charge.New(chargeParams)
}
