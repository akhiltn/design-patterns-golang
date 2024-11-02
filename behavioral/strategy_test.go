package behavioral

import "testing"

func TestStrategy(t *testing.T) {
	testMap := map[string]PaymentGateway{
		"creditcard": &CreditcardPayment{},
		"paypal":     &PaypalPayment{},
		"bitcoin":    &BitcoinPayment{},
	}
	for key, value := range testMap {
		cart := &ShoppingCart{value}
		if key != cart.payment.Pay(100) {
			t.Errorf("Expected %s but got %s", key, cart.payment.Pay(100))
		}
	}
}
