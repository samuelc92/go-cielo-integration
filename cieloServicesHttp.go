package cielo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

// CreditCardPaymentTransactionPost make post requestion on Cielo
func CreditCardPaymentTransactionPost(r *requestCreditCard, url, merchantId,
	merchantKey string) *Response {

	//Build body of request
	body, err := json.Marshal(r)
	if err != nil {
		return &Response{
			ReturnMessage: err.Error(),
			Success:       false}
	}

	//Build and do post request
	req, err := http.NewRequest(http.MethodPost, url+"1/sales",
		bytes.NewBuffer(body))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "aplication/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(body)))
	req.Header.Add("MerchantId", merchantId)
	req.Header.Add("MerchantKey", merchantKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &Response{
			ReturnMessage: err.Error(),
			Success:       false}
	}

	defer resp.Body.Close()

	response := &Response{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return &Response{
			ReturnMessage: err.Error(),
			Success:       false}
	}

	return response
}
