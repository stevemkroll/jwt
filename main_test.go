package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

const wantHeader string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
const wantPayload string = "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"
const wantSignature string = "z44tlyeOKLLrGMdctidcC7kZ6i8jQ4LWv1UogjXSnlI"

const wantJWT string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.z44tlyeOKLLrGMdctidcC7kZ6i8jQ4LWv1UogjXSnlI"

func TestHeader(t *testing.T) {
	h := []byte(`{"alg":"HS256","typ":"JWT"}`)
	header := base64.RawURLEncoding.EncodeToString(h)
	if header != wantHeader {
		t.Fatal("header does not match")
	}
}

func TestPayload(t *testing.T) {
	p := []byte(`{"sub":"1234567890","name":"John Doe","iat":1516239022}`)
	payload := base64.RawURLEncoding.EncodeToString(p)
	if payload != wantPayload {
		t.Fatal("payload does not match")
	}
}

func TestSignature(t *testing.T) {
	s := hmac.New(sha256.New, []byte("hello_world"))
	s.Write([]byte(fmt.Sprintf("%+v.%+v", wantHeader, wantPayload)))
	signature := base64.RawURLEncoding.EncodeToString(s.Sum(nil))
	if signature != wantSignature {
		t.Fatal("signature does not match")
	}
}

func TestJWT(t *testing.T) {
	h := []byte(`{"alg":"HS256","typ":"JWT"}`)
	header := base64.RawURLEncoding.EncodeToString(h)
	p := []byte(`{"sub":"1234567890","name":"John Doe","iat":1516239022}`)
	payload := base64.RawURLEncoding.EncodeToString(p)
	s := hmac.New(sha256.New, []byte("hello_world"))
	s.Write([]byte(fmt.Sprintf("%+v.%+v", header, payload)))
	signature := base64.RawURLEncoding.EncodeToString(s.Sum(nil))
	jwt := fmt.Sprintf("%+v.%+v.%+v", header, payload, signature)
	if jwt != wantJWT {
		t.Fatal("jwt does not match")
	}
	t.Log(jwt)
}
