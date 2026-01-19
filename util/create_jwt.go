package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"isShopOwner"`
}

func CreateJwt(data Payload, secret string) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArrHeader, err := json.Marshal(header)
	if(err != nil) {
		return "", err
	}
	headerB64 := Base64UrlEncoder(byteArrHeader)
	
	byteArrPayload, err := json.Marshal(data)
	if(err != nil) {
		return "", err
	}
	PayloadB64 := Base64UrlEncoder(byteArrPayload)

	message := headerB64 + "." + PayloadB64
	signature := hmac.New(sha256.New, []byte(secret)).Sum([]byte(message))
	signatureB64 := Base64UrlEncoder(signature)
	
	token := headerB64 + "." + PayloadB64 + "." + signatureB64
	return token, nil
}

func Base64UrlEncoder(byteArr []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(byteArr)
} 