package whmcs

import (
	"errors"
	"strings"
)

// Errors which are returned when validating orders for completeness/accuracy.
var (
	ErrNoClientID        = errors.New("Empty client ID")
	ErrNoPID             = errors.New("Empty PID")
	ErrNoDomain          = errors.New("Empty domain")
	ErrNoBillingCycle    = errors.New("Empty billing cycle")
	ErrNoDomainType      = errors.New("Empty domain type")
	ErrInvalidDomainType = errors.New("Invalid domain type")
	ErrNoRegPeriod       = errors.New("Empty registration period")
	ErrNoNameserver      = errors.New("Empty nameserver")
	ErrNoPaymentMethod   = errors.New("Empty payment method")
)

// Order is all the data needed to create a WHMCS order.
type Order struct {

	// Required attributes
	ClientID      int64  `json:"clientid,string"`
	PID           int64  `json:"pid,string,omitempty"`
	Domain        string `json:"domain"`
	BillingCycle  string `json:"billingcycle"`
	DomainType    string `json:"domaintype"`
	Regperiod     int64  `json:"regperiod,string"`
	EppCode       string `json:"eppcode"`
	Nameserver1   string `json:"nameserver1"`
	PaymentMethod string `json:"paymentmethod"`

	// Optional attributes
	CustomFields   CustomFields  `json:"customfields,omitempty"`
	ConfigOptions  ConfigOptions `json:"configoptions,omitempty"`
	PriceOverride  float64       `json:"priceoverride,string,omitempty"`
	PromoCode      string        `json:"promocode,omitempty"`
	PromoOverride  string        `json:"promooverride,omitempty"`
	AffID          string        `json:"affid,omitempty"`
	NoInvoice      bool          `json:"noinvoice,string,omitempty"`
	NoInvoiceEmail bool          `json:"noinvoiceemail,string,omitempty"`
	ClientIP       string        `json:"clientip,omitempty"`
	Addons         string        `json:"addons,omitempty"`

	// For VPS/Dedicated Server Orders only
	Hostname  string `json:"hostname,omitempty"`
	NS1Prefix string `json:"ns1prefix,omitempty"`
	NS2Prefix string `json:"ns2prefix,omitempty"`
	RootPw    string `json:"rootpw,omitempty"`

	// For domain reg only
	ContactID           int64  `json:"contactid,omitempty"`
	DNSManagement       bool   `json:"dnsmanagement,omitempty"`
	EmailForwarding     bool   `json:"emailforwarding,omitempty"`
	IDProtection        bool   `json:"idprotection,omitempty"`
	Nameserver2         string `json:"nameserver2,omitempty"`
	Nameserver3         string `json:"nameserver3,omitempty"`
	Nameserver4         string `json:"nameserver4,omitempty"`
	DomainRenewOverride int64  `json:"domainrenewoverride,omitempty"`
	DomainPriceOverride int64  `json:"domainpriceoverride,omitempty"`
}

func (o *Order) Error() error {
	return nil
}

// OrderResponse holds the successful order information sent via from the WHMCS API.
type OrderResponse struct {
	Result     string `json:"result"`
	OrderID    int64  `json:"orderid"`
	InvoiceID  int64  `json:"invoiceid"`
	ProductIDs string `json:"productids"`
	AddonIDs   string `json:"addonids"`
	DomainIDs  string `json:"domainids"`
}

// Products returns a slice of product ids related to the order.
func (o *OrderResponse) Products() []string {
	return strings.Split(o.ProductIDs, "/")
}

// Addons returns a slice of addon ids related to the order.
func (o *OrderResponse) Addons() []string {
	return strings.Split(o.AddonIDs, ",")
}

// Domains returns a slice of domains related to the order.
func (o *OrderResponse) Domains() []string {
	return strings.Split(o.DomainIDs, ",")
}

// AcceptOrderRequest is a struct of available parameters to accept an order.
type AcceptOrderRequest struct {
	OrderID int64 `json:"orderid,string"`

	// Optional attributes
	ServerID        int64  `json:"serverid,string,omitempty"`
	ServiceUsername string `json:"serviceusername,omitempty"`
	ServicePassword string `json:"servicepassword,omitempty"`
	AutoSetup       bool   `json:"autosetup,string,omitempty"`
	Registrar       string `json:"registrar,omitempty"`
	SendRegistrar   bool   `json:"sendregistrar,string,omitempty"`
	SendEmail       bool   `json:"sendemail,string,omitempty"`
}

func (o *AcceptOrderRequest) Error() error {
	return nil
}
