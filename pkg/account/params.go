package account

// FetchAccountParams represents parameters to pass when fetching an account.
type FetchAccountParams struct {
	// ID represents the account ID to be fetched.
	ID string
}

// DeleteAccountParams represents parameters to pass when deleting an account.
type DeleteAccountParams struct {
	// ID represents the account ID to be deleted.
	ID string

	// Version represents the account version that should be deleted.
	Version int64
}
