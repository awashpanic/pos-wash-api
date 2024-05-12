package constant

type PaymentChannel string

const (
	BankPaymentChannel    PaymentChannel = "bank-transfer"
	CashPaymentChannel    PaymentChannel = "cash"
	EwalletPaymentChannel PaymentChannel = "ewallet"
	QrisPaymentChannel    PaymentChannel = "qris"
)
