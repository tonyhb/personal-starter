package auth

import (
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

// MakeJWT creates a JWT to authroize a given user ID for the specific
// brand of KeepUpdated (as specified by the "audience" string).
//
// The audience string is the domain name in which the user logged in.
func MakeJWT(userID, audience string, expiry time.Time) ([]byte, error) {
	claims := jws.Claims{}
	claims.SetExpiration(expiry)
	claims.SetIssuedAt(time.Now().Add(-1 * time.Minute))
	claims.SetSubject(userID)
	claims.SetAudience(audience)

	token := jws.NewJWT(claims, crypto.SigningMethodHS512)
	// TODO: replace password here
	byt, err := token.Serialize([]byte("replace-with-password-from-vault-21637wdsdb"))
	return byt, err
}
