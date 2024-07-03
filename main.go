package strategy

func main() {
	product := "vehicle"
	payWay := 2

	var payment Payment

	switch payWay {
	case 1:
		payment = NewCashPayment()
	case 2:
		payment = NewCreditCardPayment("8888444422220000", "123")
	case 3:
		payment = NewPayPalPayment()
	case 4:
		payment = NewQiwiPayment()
	default:
		panic("Invalid payment method")
	}

	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	// implementation

	err := payment.Pay()
	if err != nil {
		// handle payment error
		return
	}
}

type Payment interface {
	Pay() error
}

type cashPayment struct {
	// implementation
}

func NewCashPayment() Payment {
	return nil
}

func (p *cashPayment) Pay() error {
	// implementation
	return nil
}

type creditCardPayment struct {
	// implementation
	cardNumber, cvv string
}

func NewCreditCardPayment(cardNumber, cvv string) Payment {
	return &creditCardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *creditCardPayment) Pay() error {
	// implementation
	return nil
}

type payPalPayment struct {
	// implementation
}

func NewPayPalPayment() Payment {
	return nil
}

func (p *payPalPayment) Pay() error {
	// implementation
	return nil
}

type qiwiPayment struct {
	// implementation
}

func NewQiwiPayment() Payment {
	return nil
}

func (p *qiwiPayment) Pay() error {
	// implementation
	return nil
}
