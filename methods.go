package cielo

import (
	"errors"
	"strconv"
	"strings"
)

//CONSTRUCTORS

// NewRequestCreditCard constructor of Request Credit Card
func NewRequestCreditCard(merchantOrderId string, payment payment) *requestCreditCard {
	return &requestCreditCard{
		MerchantOrderId: merchantOrderId,
		Payment:         payment}
}

// NewCompleteRequestCreditCard constructor of Request Credit Card
func NewCompleteRequestCreditCard(merchantOrderId string, payment payment, customer *customer) *requestCreditCard {
	return &requestCreditCard{
		MerchantOrderId: merchantOrderId,
		Customer:        customer,
		Payment:         payment}
}

//NewCreditCard constructor of Credit Card class
func NewCreditCard(cardNumber, holder, expirationDate, securityCode, brand string) *creditCard {
	return &creditCard{
		CardNumber:     cardNumber,
		Holder:         holder,
		ExpirationDate: expirationDate,
		SecurityCode:   securityCode,
		Brand:          brand,
	}
}

//NewCreditPayment is the constructor of Credit Payment
func NewCreditPayment(types, softDescriptor string, amount, installments int,
	c *creditCard) *creditPayment {
	return &creditPayment{
		Type:           types,
		SoftDescriptor: softDescriptor,
		Amount:         amount,
		Installments:   installments,
		CreditCard:     c,
	}
}

//NewCieloIntegration30 constructor of Cielo Integration 3.0 class
func NewCieloIntegration30(env EnvironmentEnum, merchantId, merchantKey string) *cieloIntegration30 {
	c := &cieloIntegration30{
		Environment: env,
		MerchantId:  merchantId,
		MerchantKey: merchantKey}
	c.returnUrl()
	return c
}

// ReturnResponse return the response of cielo
func (c *cieloIntegration30) ReturnResponse(paymentID, tid, status, returnCode, returnMessage string,
	success bool) *Response {
	return &Response{
		PaymentId:     paymentID,
		Tid:           tid,
		ReturnCode:    returnCode,
		ReturnMessage: returnMessage,
	}
}

// VALIDATE

//ValidatePayment valid the attributes of creditPayment
func (p *creditPayment) ValidatePayment() (bool, error) {
	if p.Type == "" {
		return false, errors.New("Type is obligation")
	}

	if len(p.Type) > 100 {
		return false, errors.New("Type can't has more than 100 characters")
	}

	if p.SoftDescriptor == "" {
		return false, errors.New("SoftDescriptor is obligation")
	}

	if len(p.SoftDescriptor) > 13 {
		return false, errors.New("SoftDescriptor can't has more than 13 characters")
	}

	return true, nil
}

//ValidatePayment valid the attributes of creditPayment
func (r *requestCreditCard) ValidateRequest() (bool, error) {
	if r.MerchantOrderId == "" {
		return false, errors.New("MerchantOrderId is obligation")
	}

	if len(r.MerchantOrderId) > 50 {
		return false, errors.New("MerchantOrderId is invalid")
	}

	return true, nil
}

//ValidateCard valid attributes of Credit Card class
func (c *creditCard) ValidateCard() (bool, error) {
	if c.CardNumber == "" {
		return false, errors.New("Card number is obligation")
	}

	if len(c.CardNumber) > 16 {
		return false, errors.New("Card number can't has more than 16 characters")
	}

	if c.Holder == "" {
		return false, errors.New("Holder is obligation")
	}

	if len(c.Holder) > 25 {
		return false, errors.New("Holder can't has more than 25 characters")
	}

	if c.ExpirationDate == "" {
		return false, errors.New("Expiration date is obligation")
	}

	if len(c.ExpirationDate) > 7 {
		return false, errors.New("Expiration date can't has more than 7 characters")
	}

	arr := strings.Split(c.ExpirationDate, "/")
	month, err := strconv.Atoi(arr[0])

	if err != nil {
		return false, errors.New("Expiration date is invalid")
	}

	if month < 1 || month > 12 {
		return false, errors.New("Expiration date is invalid")
	}

	if c.Brand == "" {
		return false, errors.New("Brand is obligation")
	}

	if len(c.Brand) > 10 {
		return false, errors.New("Brand can't has more than 10 characters")
	}

	return true, nil
}

//Validate valid the attributes of cieloIntegration30 class
func (c *cieloIntegration30) Validate() (bool, error) {
	if c.MerchantId == "" {
		return false, errors.New("MerchantId is obligation")
	}

	if len(c.MerchantId) != 36 {
		return false, errors.New("MerchantId is invalid")
	}

	if c.MerchantKey == "" {
		return false, errors.New("MerchantKey is obligation")
	}

	if len(c.MerchantKey) != 40 {
		return false, errors.New("MerchantKey is invalid")
	}

	return true, nil
}

// Methods

// CreditCardPaymentTransaction realize simple transaction of credit card payments in cielo
func (c *cieloIntegration30) CreditCardPaymentTransaction(r *requestCreditCard) *Response {
	return CreditCardPaymentTransactionPost(r, c.TranactionURL, c.MerchantId, c.MerchantKey)
}

//returnUrl return API's url
func (c *cieloIntegration30) returnUrl() {
	switch c.Environment {
	case TEST:
		c.TranactionURL = "https://apisandbox.cieloecommerce.cielo.com.br/"
		c.ConsultURL = "https://apiquerysandbox.cieloecommerce.cielo.com.br/"
	case PRODUCTION:
		c.TranactionURL = "https://api.cieloecommerce.cielo.com.br/"
		c.ConsultURL = "https://apiquery.cieloecommerce.cielo.com.br/"
	default:
		c.TranactionURL = ""
		c.ConsultURL = ""
	}
}
