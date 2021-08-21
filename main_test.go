package jwt_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"testing"
)

const (
	header  = `{"alg":"HS256","typ":"JWT"}`
	payload = `{"sub":"1234567890","name":"John Doe","iat":1516239022}`
	secret  = `your-256-bit-secret`
)

const (
	result                = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`
	resultBase64Signature = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.cThIIoDvwdueQB468K5xDc5633seEFoqwxjF_xSJyQQ`
)

var (
	encodedHeader    string
	encodedPayload   string
	encodedSignature string
)

var (
	ErrHeaderMatch    error = errors.New("err header does not match")
	ErrPayloadMatch   error = errors.New("err payload does not match")
	ErrSignatureMatch error = errors.New("err signature does not match")
	ErrTokenMatch     error = errors.New("err token does not match")
)

func TestHeader(t *testing.T) {
	encodedHeader = base64.RawURLEncoding.EncodeToString([]byte(header))
	if encodedHeader != strings.Split(result, ".")[0] {
		t.Fatal(ErrHeaderMatch)
	}
}

func TestPayload(t *testing.T) {
	encodedPayload = base64.RawURLEncoding.EncodeToString([]byte(payload))
	if encodedPayload != strings.Split(result, ".")[1] {
		t.Fatal(ErrPayloadMatch)
	}
}

func TestSignature(t *testing.T) {
	TestHeader(t)
	TestPayload(t)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(encodedHeader + "." + encodedPayload)); err != nil {
		t.Fatal(err)
	}
	encodedSignature = base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	if encodedSignature != strings.Split(result, ".")[2] {
		t.Fatal(ErrSignatureMatch)
	}
}

func TestBase64Signature(t *testing.T) {
	TestHeader(t)
	TestPayload(t)
	b, err := base64.RawURLEncoding.DecodeString(secret)
	if err != nil {
		t.Fatal(err)
	}
	h := hmac.New(sha256.New, b)
	if _, err := h.Write([]byte(encodedHeader + "." + encodedPayload)); err != nil {
		t.Fatal(err)
	}
	encodedSignature = base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	if encodedSignature != strings.Split(resultBase64Signature, ".")[2] {
		t.Fatal(ErrSignatureMatch)
	}
}

func TestToken(t *testing.T) {
	TestHeader(t)
	TestPayload(t)
	TestSignature(t)
	token := fmt.Sprintf("%s.%s.%s", encodedHeader, encodedPayload, encodedSignature)
	if token != result {
		t.Fatal(ErrTokenMatch)
	}
}

func TestTokenBase64Signature(t *testing.T) {
	TestHeader(t)
	TestPayload(t)
	TestBase64Signature(t)
	token := fmt.Sprintf("%s.%s.%s", encodedHeader, encodedPayload, encodedSignature)
	if token != resultBase64Signature {
		t.Fatal(ErrTokenMatch)
	}
}
