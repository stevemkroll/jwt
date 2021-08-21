package header

import (
	"errors"
	"jwt/alg"
	"jwt/typ"
	"testing"
)

var h Header
var ErrHeaderDoesNotMatch error = errors.New("err header does not match")

const expected string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
const expectedJSON string = `{"alg":"HS256","typ":"JWT"}`

func TestEncodeToString(t *testing.T) {
	// set header values
	h.Alg = alg.HS256
	h.Typ = typ.JWT

	// encode to string
	res, err := h.EncodeToString()
	if err != nil {
		t.Fatal(err)
	}

	// validate result
	if *res != expected {
		t.Fatal()
	}
}

func TestDecodeString(t *testing.T) {
	// decode to string
	res, err := DecodeToString(expected)
	if err != nil {
		t.Fatal(err)
	}

	// validate result
	if *res != expectedJSON {
		t.Fatal()
	}
}
