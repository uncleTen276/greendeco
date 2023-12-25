package models

import "github.com/google/uuid"

const VnPayUrl = "https://sandbox.vnpayment.vn/paymentv2/vpcpay.html?"

type PaymentRequest struct {
	Id   uuid.UUID `json:"id" db:"id"`
	Type string    `json:"type"`
}

type PayPalReturn struct {
	ID string `json:"id"`
}

type PaymentCurrenctResponse struct {
	Meta struct {
		Code       int    `json:"code"`
		Disclaimer string `json:"disclaimer"`
	} `json:"meta"`
	Response struct {
		Timestamp int     `json:"timestamp"`
		Date      string  `json:"date"`
		From      string  `json:"from"`
		To        string  `json:"to"`
		Amount    int     `json:"amount"`
		Value     float64 `json:"value"`
	} `json:"response"`
	Timestamp int     `json:"timestamp"`
	Date      string  `json:"date"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    int     `json:"amount"`
	Value     float64 `json:"value"`
}
