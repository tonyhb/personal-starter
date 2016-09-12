package v0

import (
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/tonyhb/keepupdated/pkg/api/apilib"
	"gitlab.com/tonyhb/keepupdated/pkg/api/auth"
	"gitlab.com/tonyhb/keepupdated/pkg/api/v0/errors"
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
			Status: http.StatusCreated,
			Data:   responses.JWT{},
		},
	}
}

func (v *V0) Login(ctx context.Context, req *restful.Request, w *restful.Response) interface{} {
	data := new(forms.EmailPassAuth)
	if err := req.ReadEntity(data); err != nil {
		return v.WrapError(err)
	}

	u, err := v.mgr.UserByEmail(data.Email)
	if err != nil {
		return errors.ErrInvalidCredentials
	}
	err = u.CheckPassword(data.Password)
	if err != nil {
		return errors.ErrInvalidCredentials
	}

	jwt, _ := auth.MakeJWT(strconv.Itoa(u.ID), "keepupdated.com", time.Now().Add(24*time.Hour))
	return responses.MakeJWT(jwt)
}

func (v *V0) Register(ctx context.Context, req *restful.Request, w *restful.Response) interface{} {
	register := new(forms.Register)
	if err := req.ReadEntity(register); err != nil {
		return v.WrapError(err)
	}
	if err := validate.Run(register); err != nil {
		// TODO: make a function in the errors package to wrap and format
		// validation errors from govalidate
		return nil
	}

	// create a new account and user
	acct := register.Account()
	user := register.User()

	if err := v.mgr.CreateAccount(acct); err != nil {
		return v.WrapError(err)
	}
	user.AccountID = acct.ID
	if err := v.mgr.CreateUser(user); err != nil {
		return v.WrapError(err)
	}

	jwt, _ := auth.MakeJWT(strconv.Itoa(user.ID), "keepupdated.com", time.Now().Add(24*time.Hour))

	// TODO: kafka job to send a welcome email
	return responses.Register{
		User:    responses.MakeUser(*user),
		Account: responses.MakeAccount(*acct),
		JWT:     responses.MakeJWT(jwt),
	}
}
