package account

// Attributes represents the domain model for account attributes.
type Attributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 string   `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
	CustomerID              *string  `json:"customer_id,omitempty"`
	Name                    []string `json:"name,omitempty"`
}

// NewAttributes returns an account attribute builder.
//
// baseCurrency and country are both required to create a new account.
func NewAttributes(baseCurrency string, country string) *Attributes {
	return &Attributes{
		BaseCurrency: baseCurrency,
		Country:      country,
	}
}

// WithAccountClassification sets the account classification.
func (a *Attributes) WithAccountClassification(class *string) *Attributes {
	a.AccountClassification = class
	return a
}

// WithAccountMatchingOptOut sets the opt-out for the account.
func (a *Attributes) WithAccountMatchingOptOut(opt *bool) *Attributes {
	a.AccountMatchingOptOut = opt
	return a
}

// WithAccountNumber sets the account number for the account.
func (a *Attributes) WithAccountNumber(num string) *Attributes {
	a.AccountNumber = num
	return a
}

// WithBankID sets the bank ID for the account.
func (a *Attributes) WithBankID(id string) *Attributes {
	a.BankID = id
	return a
}

// WithBankIDCode sets the bank id code for the bank.
func (a *Attributes) WithBankIDCode(code string) *Attributes {
	a.BankIDCode = code
	return a
}

// WithBaseCurrency sets the base currency for the account.
func (a *Attributes) WithBaseCurrency(curr string) *Attributes {
	a.BaseCurrency = curr
	return a
}

// WithBic sets the BIC number for the account.
func (a *Attributes) WithBic(bic string) *Attributes {
	a.Bic = bic
	return a
}

// WithCountry sets the country for the account.
func (a *Attributes) WithCountry(country string) *Attributes {
	a.Country = country
	return a
}

// WithIban sets the IBAN number for the account.
func (a *Attributes) WithIban(iban string) *Attributes {
	a.Iban = iban
	return a
}

// WithJointAccount indicates that the account is a joint account.
func (a *Attributes) WithJointAccount(isJoint *bool) *Attributes {
	a.JointAccount = isJoint
	return a
}

// WithSecondaryIdentification sets the secondary identification for the account.
func (a *Attributes) WithSecondaryIdentification(sec string) *Attributes {
	a.SecondaryIdentification = sec
	return a
}

// WithStatus sets the account status.
func (a *Attributes) WithStatus(status *string) *Attributes {
	a.Status = status
	return a
}

// WithSwitched sets if the account has switched or not.
func (a *Attributes) WithSwitched(isSwitched *bool) *Attributes {
	a.Switched = isSwitched
	return a
}

// WithCustomerID sets the customer ID.
func (a *Attributes) WithCustomerID(customerID *string) *Attributes {
	a.CustomerID = customerID
	return a
}

// WithName sets the name.
func (a *Attributes) WithName(name string) *Attributes {
	a.Name = append(a.Name, name)
	return a
}
