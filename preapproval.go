package gopayhere

type PreapprovalNotification struct {
	MerchantID     string `json:"merchant_id" schema:"merchant_id"`
	OrderID        string `json:"order_id" schema:"order_id"`
	PaymentID      string `json:"payment_id" schema:"payment_id"`
	Amount         string `json:"amount" schema:"payhere_amount"`
	Currency       string `json:"currency" schema:"payhere_currency"`
	StatusCode     int    `json:"status_code" schema:"status_code"`
	Md5sig         string `json:"md5sig" schema:"md5sig"`
	StatusMessage  string `json:"status_message" schema:"status_message"`
	CustomerToken  string `json:"customer_token" schema:"customer_token"`
	Custom1        string `json:"custom_1" schema:"custom_1"`
	Custom2        string `json:"custom_2" schema:"custom_2"`
	Method         string `json:"method" schema:"method"`
	CardHolderName string `json:"card_holder_name" schema:"card_holder_name"`
	CardNumber     string `json:"card_no" schema:"card_no"`
	CardExpiry     string `json:"card_expiry" schema:"card_expiry"`
}
