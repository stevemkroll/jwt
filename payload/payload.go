package payload

import "time"

type Payload struct {
	Iss string    `json:"iss,omitempty"` // issuer
	Sub string    `json:"sub,omitempty"` // subject
	Aud string    `json:"aud,omitempty"` // audience
	Exp time.Time `json:"exp,omitempty"` // expiration (time)
	Nbf time.Time `json:"nbf,omitempty"` // not before (time)
	Iat time.Time `json:"iat,omitempty"` // issued at (time)
	Jti string    `json:"jti,omitempty"` // JWT ID
}
