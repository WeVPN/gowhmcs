package whmcs

import (
	"errors"
)

var (
	ErrTransactionAmountsEmpty       = errors.New("Transaction amounts empty")
	ErrTransactionPaymentMethodEmpty = errors.New("Transaction payment method empty")
	ErrTransactionDateEmpty          = errors.New("Transaction date empty")
)

type Transaction struct {
	// Required fields
	AmountIn      float64 `json:"amountin,string" xml:"amountin"`
	AmountOut     float64 `json:"amountout,string" xml:"amountout"`
	PaymentMethod string  `json:"paymentmethod" xml:"paymentmethod"`
	Date          Date    `json:"date" xml:"date"`

	// Optional fields
	UserID      string  `json:"userid,omitempty" xml:"userid,omitempty"`
	InvoiceID   string  `json:"invoiceid,omitempty" xml:"invoiceid,omitempty"`
	Description string  `json:"description,omitempty" xml:"description"`
	Fees        float64 `json:"fees,string",omitempty xml:"fees"`
	TransID     string  `json:"transid,omitempty" xml:"transid"`
	Credit      bool    `json:"credit,string" xml:"credit"`
}

func (t *Transaction) Error() error {
	if t.AmountIn == 0 && t.AmountOut == 0 {
		return ErrTransactionAmountsEmpty
	}
	if t.PaymentMethod == "" {
		return ErrTransactionPaymentMethodEmpty
	}
	if t.Date.Time().IsZero() {
		return ErrTransactionDateEmpty
	}
	return nil
}

func (a *API) AddTransaction(t *Transaction) error {

	if err := t.Error(); err != nil {
		return err
	}

	if _, err := a.Do("addtransaction", &t); err != nil {
		return err
	}

	return nil

}
