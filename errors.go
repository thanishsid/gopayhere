package gopayhere

import "errors"

var (
	ErrUnauthorized         = errors.New("unauthorized")
	ErrInvalidAccessToken   = errors.New("invalid access token")
	ErrAccessTokenExpired   = errors.New("access token has expired")
	ErrInvalidCustomerToken = errors.New("invalid customer token")

	ErrInvalidCurrency = errors.New("invalid currency")
	ErrInvalidAmount   = errors.New("invalid amount")

	ErrCancelled = errors.New("operation cancelled")
	ErrFailed    = errors.New("operation failed")
	ErrNoData    = errors.New("data not found")
)
