package jwt

type Header struct {
	Algorithm algorithm `json:"alg"`
	Token     token     `json:"typ"`
}
