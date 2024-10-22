package behavioral

type PaymentGateway interface {
	Pay(amount float64) string
}

type CreditcardPayment struct{}

func (p *CreditcardPayment) Pay(amount float64) string {
	return "creditcard"
}

type PaypalPayment struct{}

func (p *PaypalPayment) Pay(amount float64) string {
	return "paypal"
}

type BitcoinPayment struct{}

func (p *BitcoinPayment) Pay(amount float64) string {
	return "bitcoin"
}

type ShoppingCart struct {
	payment PaymentGateway
}
