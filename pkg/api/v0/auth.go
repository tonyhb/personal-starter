package v0

import (
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/tonyhb/keepupdated/pkg/api/apilib"
	"gitlab.com/tonyhb/keepupdated/pkg/api/v0/forms"
	"gitlab.com/tonyhb/keepupdated/pkg/api/v0/responses"

	"github.com/emicklei/go-restful"
	"github.com/tonyhb/govalidate"
)

func (v *V0) LoginRoute() apilib.Route {
	return apilib.Route{
		Path:    "/login",
		Method:  "POST",
		Handler: v.Login,
		Returns: apilib.Returns{
			Status: http.StatusOK,
			Data:   responses.JWT{},
		},
	}
}

func (v *V0) RegisterRoute() apilib.Route {
	return apilib.Route{
		Path:    "/register",
		Method:  "POST",
		Handler: v.Register,
		Returns: apilib.Returns{
			Status: http.StatusOK,
			Data:   responses.JWT{},
		},
	}
}

func (v *V0) Login(ctx context.Context, req *restful.Request, w *restful.Response) {
	// TODO: auth the user via username/password
	w.Write([]byte("lol"))
}

func (v *V0) Register(ctx context.Context, req *restful.Request, w *restful.Response) {
	register := new(forms.Register)
	if err := req.ReadEntity(register); err != nil {
		return
	}
	if err := validate.Run(register); err != nil {
		return
	}

	// create a new account and user
	acct := register.Account()
	user := register.User()

	if err := v.mgr.CreateAccount(acct); err != nil {
		return
	}
	user.AccountID = acct.ID
	if err := v.mgr.CreateUser(user); err != nil {
		return
	}

	jwt, _ := apilib.MakeJWT(strconv.Itoa(user.ID), "keepupdated.com", time.Now().Add(24*time.Hour))

	// TODO: kafka job to send a welcome email
	// TODO: create JWT and log in user
	w.WriteEntity(responses.Register{
		User:    responses.MakeUser(*user),
		Account: responses.MakeAccount(*acct),
		JWT:     responses.MakeJWT(jwt),
	})
}
