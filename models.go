package cielo

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

type requestCreditCard struct {
	MerchantOrderId string
	Customer        *customer `json:",omitempty"`
	Payment         payment
}

type creditPayment struct {
	Type                 string
	Amount, Installments int
	SoftDescriptor       string
	CreditCard           *creditCard
}

type creditCard struct {
	CardNumber, Holder, ExpirationDate, SecurityCode, Brand string
}

// Response return of cielo integration
type Response struct {
	PaymentId, Tid, Status, ReturnCode, ReturnMessage string
	Success                                           bool
}

// CieloIntegration30 realize cielo integration version 3.0
type cieloIntegration30 struct {
	TranactionURL, ConsultURL, MerchantId, MerchantKey string
	Environment                                        EnvironmentEnum
}

type customer struct {
	Name string `json:",omitempty"`
}
