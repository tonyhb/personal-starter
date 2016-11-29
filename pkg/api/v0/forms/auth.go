package forms

import (
	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

type Register struct {
	StripeID string `json:"stripeId" validate:"NotEmpty"`
	Plan     string `json:"plan"     validate:"NotEmpty"`
	Email    string `json:"email"    validate:"NotEmpty"`
	Password string `json:"password" validate:"NotEmpty"`
}

func (r Register) Account() *types.Account {
	acct := &types.Account{
		StripeID: r.StripeID,
	}

	switch r.Plan {
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
