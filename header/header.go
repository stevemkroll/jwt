package header

import (
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`           // algorithm
	Typ string `json:"typ,omitempty"` // media type
	Cty string `json:"cty,omitempty"` // content type
}

func (h *Header) EncodeToString() (*string, error) {
	b, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}
	res := base64.RawURLEncoding.EncodeToString(b)
	return &res, nil
}

func DecodeToString(s string) (*string, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	res := string(b)
	return &res, nil
}
