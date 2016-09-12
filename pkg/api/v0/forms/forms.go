package forms

import (
	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

type Register struct {
	StripeID string `validate:"NotEmpty"`
	Plan     string `validate:"NotEmpty"`
	Email    string `validate:"NotEmpty"`
	Password string `validate:"NotEmpty"`
}

func (r Register) Account() *types.Account {
	acct := &types.Account{
		StripeID: r.StripeID,
	}

	switch r.Plan {
	case "free":
		acct.Features = types.Free
	case "starter":
		acct.Features = types.Starter
	case "pro":
		acct.Features = types.Pro
	case "agency":
		acct.Features = types.Agency
	default:
		acct.Features = types.Free
	}

	return acct
}

func (r Register) User() *types.User {
	// Registering users always makes them a billing contact for their account
	user := &types.User{
		Email:            r.Email,
		IsBillingContact: true,
	}
	user.SetPassword(r.Password)
	return user
}

type EmailPassAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
