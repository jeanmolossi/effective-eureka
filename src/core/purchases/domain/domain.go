package domain

import "time"

// Status is the status of a purchase.
//
// Available statuses are:
// 	- APPROVED
// 	- BLOCKED
// 	- CANCELLED
// 	- CHARGEBACK
// 	- COMPLETE
// 	- EXPIRED
// 	- NO_FUNDS
// 	- OVERDUE
// 	- PARTIALLY_REFUNDED
// 	- PRE_ORDER
// 	- PRINTED_BILLET
// 	- PROCESSING_TRANSACTION
// 	- PROTESTED
// 	- REFUNDED
// 	- STARTED
// 	- UNDER_ANALISYS
// 	- WAITING_PAYMENT
type Status string

const (
	// APPROVED: The purchase was approved. It was paid and the product was sent to the buyer.
	APPROVED Status = "APPROVED"
	// BLOCKED:
	//
	// Not documented: https://help.hotmart.com/pt-BR/article/Quais-status-uma-transa%C3%A7%C3%A3o-pode-assumir-/216441297
	BLOCKED Status = "BLOCKED"
	// CANCELLED: The purchase was cancelled. When its happens look into:
	//
	// https://suportehotmart.zendesk.com/hc/pt-br/articles/203456160-Como-funciona-o-processo-de-compra-na-Hotmart-Por-que-minha-compra-foi-cancelada-
	CANCELLED Status = "CANCELLED"
	// CHARGEBACK: When client requests a chargeback because does not identify the product.
	//
	// Better UX should be custom billing to identify the product.
	CHARGEBACK Status = "CHARGEBACK"
	// COMPLETE: The purchase was completed. It means the guarantee and chargeback was expired.
	COMPLETE Status = "COMPLETE"
	// EXPIRED: Buyers that printed billet but did not pay.
	// Or else when deposit was not received after 48 hours by the bank.
	//
	// You can recover that on: https://help.hotmart.com/pt-BR/article/O-que-%C3%A9-e-como-efetuar-a-Recupera%C3%A7%C3%A3o-de-Compras-/216440947
	EXPIRED Status = "EXPIRED"
	// NO_FUNDS:
	//
	// Not documented: https://help.hotmart.com/pt-BR/article/Quais-status-uma-transa%C3%A7%C3%A3o-pode-assumir-/216441297
	NO_FUNDS Status = "NO_FUNDS"
	// OVERDUE: Client was subscribed to a product and that was not paid.
	// E.g.: The product was not paid in time.
	OVERDUE Status = "OVERDUE"
	// PARTIALLY_REFUNDED:
	//
	// Not documented: https://help.hotmart.com/pt-BR/article/Quais-status-uma-transa%C3%A7%C3%A3o-pode-assumir-/216441297
	PARTIALLY_REFUNDED Status = "PARTIALLY_REFUNDED"
	// PRE_ORDER:
	//
	// Not documented: https://help.hotmart.com/pt-BR/article/Quais-status-uma-transa%C3%A7%C3%A3o-pode-assumir-/216441297
	PRE_ORDER Status = "PRE_ORDER"
	// PRINTED_BILLET: The purchase was started and still waiting for the payment
	PRINTED_BILLET Status = "PRINTED_BILLET"
	// PROCESSING_TRANSACTION:
	//
	// Not documented: https://help.hotmart.com/pt-BR/article/Quais-status-uma-transa%C3%A7%C3%A3o-pode-assumir-/216441297
	PROCESSING_TRANSACTION Status = "PROCESSING_TRANSACTION"
	// PROTESTED: when client or producer protested the purchase, but chargeback was not approved.
	PROTESTED Status = "PROTESTED"
	// REFUNDED: The purchase was refunded. It can not be appealed.
	REFUNDED Status = "REFUNDED"
	// STARTED: The purchase was started, but not paid yet.
	STARTED Status = "STARTED"
	// UNDER_ANALISYS: The purchase was started and still waiting for the payment.
	// Frequently occur when billing is credit card or paypal.
	UNDER_ANALISYS Status = "UNDER_ANALISYS"
	// WAITING_PAYMENT: when purchase process was started but not finished.
	// still waiting for payment.
	WAITING_PAYMENT Status = "WAITING_PAYMENT"
)

type Item interface {
	Buyer() Buyer
	Product() Product
	Status() Status
	Transaction() string
	WarranyExpireDate() time.Time
}

type Buyer interface {
	Name() string
	Email() string
}

type Product interface {
	ID() uint32
	Name() string
}

type GetPurchases interface {
	Get(email, transaction string) ([]Item, error)
}
