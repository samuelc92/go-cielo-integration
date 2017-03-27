package cielo

type cieloIntegration interface {
	ReturnResponse() *Response
}

type request interface {
	ValidateRequest() (bool, error)
}

type payment interface {
	ValidatePayment() (bool, error)
}

type card interface {
	ValidateCard() (bool, error)
}
