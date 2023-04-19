// Package fixtures consists of API response entities for testing.
package fixtures

import (
	"fmt"
)

// AccountsResponseAllFields returns a JSON representation of an accounts entity.
// It returns all the fields that can be set.
func AccountsResponseAllFields(
	accountID string,
	orgID string,
	country string,
	currency string,
) string {
	return fmt.Sprintf(`{
	"data": {
    	"type": "accounts",
     	"id": "%[1]s",
      	"version": 0,
       	"organisation_id": "%[2]s",
        "attributes": {
        	"country": "%[3]s",
         	"base_currency": "%[4]s",
           	"bank_id": "20041",
            "bank_id_code": "%[3]s",
            "account_number": "0500013M026",
            "customer_id": "999",
            "iban": "%[3]s1420041010050500013M02606",
            "bic": "NWBKFR42",
            "account_classification": "Personal",
            "joint_account": false,
            "account_matching_opt_out": false,
            "switched": false,
            "status": "confirmed"
        }
    },
    "links": {
    	"self": "/accounts/%[1]s",
     	"first": "/accounts?page[number]=first",
      	"last": "/accounts?page[number]=last",
       	"next": "/accounts?page[number]=next",
        "prev": "/accounts?page[number]=prev"
    }
}`, orgID, accountID, country, currency)
}

// AccountsResponseMinFields returns a JSON representation of an accounts entity.
// It returns the minimum number of fields that can be set on it.
//
// https://www.api-docs.form3.tech/api/schemes/sepa-direct-debit/accounts/accounts/fetch-an-account.
func AccountsResponseMinFields(
	accountID string,
	orgID string,
	country string,
	currency string,
) string {
	return fmt.Sprintf(`{
	"data": {
    	"type": "accounts",
     	"id": "%[1]s",
      	"version": 0,
       	"organisation_id": "%[2]s",
        "attributes": {
        	"country": "%[3]s",
         	"base_currency": "%[4]s",
            "account_number": "0500013M026",
            "iban": "%[3]s1420041010050500013M02606",
            "account_classification": "Personal",
            "joint_account": false,
            "account_matching_opt_out": false,
            "switched": false,
            "status": "confirmed"
        }
    },
    "links": {
    	"self": "/accounts/%[1]s"
    }
}`, orgID, accountID, country, currency)
}
