package payment

import (
	"errors"
	"fmt"
	"log"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
)

type PaymentClient interface {
	CreatePayment(amount float64, userID uint, orderID string) (*stripe.CheckoutSession, error)
	GetPaymentStatus(paymentID string) (*stripe.CheckoutSession, error)
}

type payment struct {
	stripeSecretKey string
	successUrl      string
	cancelUrl       string
}

func NewPaymentClient(stripeSecretKey, successUrl, cancelUrl string) PaymentClient {
	return &payment{
		stripeSecretKey: stripeSecretKey,
		successUrl:      successUrl,
		cancelUrl:       cancelUrl,
	}
}

func (p *payment) CreatePayment(amount float64, userID uint, orderID string) (*stripe.CheckoutSession, error) {
	stripe.Key = p.stripeSecretKey
	amountInCents := amount * 100

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					UnitAmount: stripe.Int64(*stripe.Int64(int64(amountInCents))),
					Currency:   stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Electronics"),
					},
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(string(p.successUrl)),
		CancelURL:  stripe.String(string(p.cancelUrl)),
	}

	params.AddMetadata("order_id", orderID)
	params.AddMetadata("user_id", fmt.Sprintf("%d", userID))

	session, err := session.New(params)
	if err != nil {
		log.Printf("payment create session error: %v", err)
		return nil, errors.New("payment create session failed")
	}

	return session, nil
}

func (p *payment) GetPaymentStatus(paymentID string) (*stripe.CheckoutSession, error) {
	stripe.Key = p.stripeSecretKey

	session, err := session.Get(paymentID, nil)
	if err != nil {
		log.Printf("payment get session error: %v", err)
		return nil, errors.New("payment get session failed")
	}

	return session, nil
}
