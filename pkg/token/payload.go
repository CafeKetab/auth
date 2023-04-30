package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	Data      []byte `json:"data"`
	ExpiredAt int64  `json:"expired_at"`
	jwt.RegisteredClaims
}

const (
	missingData         = "missing expiration key"
	missingExpiredAt    = "missing expiration key"
	tokenHasBeenExpired = "token has been expired"
)

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if len(payload.Data) == 0 {
		errStr := fmt.Sprintf("%s: %s, payload: %v", inValidToken, missingData, payload)
		return errors.New(errStr)
	}

	if payload.ExpiredAt == 0 {
		errStr := fmt.Sprintf("%s: %s, payload: %v", inValidToken, missingExpiredAt, payload)
		return errors.New(errStr)
	}

	if exp := payload.ExpiredAt; time.Unix(int64(exp), 0).Before(time.Now()) {
		errStr := fmt.Sprintf("%s: %s, ExpiredAt: %d", inValidToken, tokenHasBeenExpired, exp)
		return errors.New(errStr)
	}

	return nil
}
