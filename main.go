package jwt

type (
	token     string
	algorithm string
)

const jwt token = "JWT"

const (
	hs256 algorithm = "HS256"
	hs384 algorithm = "HS384"
	hs512 algorithm = "HS512"

	rs256 algorithm = "RS256"
	rs384 algorithm = "RS384"
	rs512 algorithm = "RS512"

	es256 algorithm = "ES256"
	es384 algorithm = "ES384"
	es512 algorithm = "ES512"

	ps256 algorithm = "PS256"
	ps384 algorithm = "PS384"
)
