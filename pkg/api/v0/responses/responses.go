package responses

import (
	"time"

	"gitlab.com/tonyhb/keepupdated/pkg/types"

	"github.com/emicklei/go-restful"
)

type JWT struct {
	// If token were []byte this would automatically be base64 encoded by
	// go-restful's WriteEntity call.
	Token string `json:"token"`
}

func MakeJWT(jwt []byte) JWT {
	return JWT{string(jwt)}
}

type Account struct {
	ID             int       `json:"id"`
	CompanyName    string    `json:"companyName"`
	CompanyAddress string    `json:"companyAddress"`
	CompanyCity    string    `json:"companyCity"`
	CompanyState   string    `json:"companyState"`
	CompanyCountry string    `json:"companyCountry"`
	CompanyZip     string    `json:"companyZip"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func MakeAccount(t types.Account) Account {
	return Account{
		ID:             t.ID,
		CompanyName:    t.CompanyName,
		CompanyAddress: t.CompanyAddress,
		CompanyCity:    t.CompanyCity,
		CompanyState:   t.CompanyState,
		CompanyCountry: t.CompanyCountry,
		CompanyZip:     t.CompanyZip,
		CreatedAt:      t.CreatedAt,
		UpdatedAt:      t.UpdatedAt,
	}
}

type User struct {
	ID               int       `json:"id"`
	AccountID        int       `json:"accountID"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	PasswordHash     []byte    `json:"-"`
	IsBillingContact bool      `json:"isBillingContact"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func MakeUser(u types.User) User {
	return User{
		ID:               u.ID,
		AccountID:        u.AccountID,
		Name:             u.Name,
		Email:            u.Email,
		IsBillingContact: u.IsBillingContact,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}

// Register is a composite type which lists the newly created account, user,
// and a valid JWT for authenticating with the app and API immediately.
type Register struct {
	User    User    `json:"user"`
	Account Account `json:"account"`
	JWT     JWT     `json:"jwt"`
}

func (r Register) Write(w *restful.Response) {
	w.WriteHeaderAndEntity(201, r)
}
