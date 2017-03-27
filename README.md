# go-cielo-integration
Golang package to integration Cielo e-Commerce WebService.

### Creating a simple transaction
```golang
//Create and validate Cielo Integration30 object
cieloIntegration := cielo.NewCieloIntegration30(cielo.TEST, "a187afc0-db2c-4bdb-8769-678f2eb03152",
		"KRFZDQWYWVRIJIQPMRDKTDRBKQQNQGJOVTFPEVVF")
cieloIntegration.Validate()
//Create and validate Credit Card object
creditCard := cielo.NewCreditCard("1234123412341231", "Teste Holder", "12/2021", "123", "Visa")
_, err := creditCard.ValidateCard()
//Create and validate Payment object
payment := cielo.NewCreditPayment("CreditCard", "123456789ABCD", 15700, 1, creditCard)
_, err = payment.ValidatePayment()
//Create and validate Request Credit Card object
request := cielo.NewRequestCreditCard("2014111703", payment)
_, err = request.ValidateRequest()
//Do transaction
response := cieloIntegration.CreditCardPaymentTransaction(request)
