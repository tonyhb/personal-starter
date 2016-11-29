package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

var (
	ErrNoJWT         = fmt.Errorf("No JWT present in request header")
	ErrInvalidJWT    = fmt.Errorf("Invalid JWT")
	ErrExpiredJWT    = fmt.Errorf("Expired JWT")
	ErrInvalidIssuer = fmt.Errorf("Invalid issuer in JWT")
)

// MakeJWT creates a JWT to authroize a given user ID for the specific
// brand of KeepUpdated (as specified by the "audience" string).
//
// The issuer string is the domain name in which the user logged in; this is
// used to ensure the domain matches each time the JWT is used.
func MakeJWT(userID, issuer string, expiry time.Time) ([]byte, error) {
	claims := jws.Claims{}
	claims.SetExpiration(expiry)
	claims.SetIssuedAt(time.Now().Add(-1 * time.Minute))
	claims.SetSubject(userID)
	claims.SetIssuer(issuer)

	token := jws.New(claims, crypto.SigningMethodHS512)
	// TODO: replace password here
	byt, err := token.Compact([]byte("replace-with-password-from-vault-21637wdsdb"))
	return byt, err
}

func ParseJWT(r *http.Request) (string, error) {
	jwt, err := jws.ParseJWTFromRequest(r)
	if err != nil {
		return "", ErrNoJWT
	}

	err = jwt.Validate([]byte("replace-with-password-from-vault-21637wdsdb"), crypto.SigningMethodHS512)
	if err != nil {
		return "", ErrInvalidJWT
	}

	err = jwt.Claims().Validate(time.Now(), 60*time.Second, 60*time.Second)
	if err != nil {
		return "", ErrExpiredJWT
	}

	id, ok := jwt.Claims().Get("sub").(string)
	if !ok {
		return "", ErrInvalidJWT
	}

	iss, ok := jwt.Claims().Get("iss").(string)
	if !ok {
		return "", ErrInvalidJWT
	}

	if r.URL.Host != iss {
		return "", ErrInvalidIssuer
	}

	return id, nil
}
