package gopayhere

import "net/url"

type PreapprovalNotification struct {
	MerchantID     string `json:"merchant_id"`
	OrderID        string `json:"order_id"`
	PaymentID      string `json:"payment_id"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	StatusCode     string `json:"status_code"`
	Md5sig         string `json:"md5sig"`
	StatusMessage  string `json:"status_message"`
	CustomerToken  string `json:"customer_token"`
	Custom1        string `json:"custom_1"`
	Custom2        string `json:"custom_2"`
	Method         string `json:"method"`
	CardHolderName string `json:"card_holder_name"`
	CardNumber     string `json:"card_no"`
	CardExpiry     string `json:"card_expiry"`
}

func GetPrepprovalNotificationFromUrlValues(f url.Values) PreapprovalNotification {
	return PreapprovalNotification{
		MerchantID:     f.Get("merchant_id"),
		OrderID:        f.Get("order_id"),
		PaymentID:      f.Get("payment_id"),
		Amount:         f.Get("payhere_amount"),
		Currency:       f.Get("payhere_currency"),
		StatusCode:     f.Get("status_code"),
		Md5sig:         f.Get("md5sig"),
		StatusMessage:  f.Get("status_message"),
		CustomerToken:  f.Get("customer_token"),
		Custom1:        f.Get("custom_1"),
		Custom2:        f.Get("custom_2"),
		Method:         f.Get("method"),
		CardHolderName: f.Get("card_holder_name"),
		CardNumber:     f.Get("card_no"),
		CardExpiry:     f.Get("card_expiry"),
	}
}
