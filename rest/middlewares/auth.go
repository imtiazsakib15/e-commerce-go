package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/util"
	"net/http"
	"strings"
)

func (m *Middlewares) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		headerAuth := r.Header.Get("Authorization")
		if(headerAuth == "") {
			http.Error(w, "Unauthorized!!!", http.StatusUnauthorized)
			return
		}

		headerAuthArr := strings.Split(headerAuth, " ")
		if(len(headerAuthArr) != 2) {
			http.Error(w, "Unauthorized!!!", http.StatusUnauthorized)
			return
		}

		accessToken := headerAuthArr[1]
		tokenArr := strings.Split(accessToken, ".")
		if(len(tokenArr) != 3) {
			http.Error(w, "Unauthorized!!!", http.StatusUnauthorized)
			return
		}

		tokenHeader := tokenArr[0]
		tokenPayload := tokenArr[1]
		tokenSignature := tokenArr[2]

		message := tokenHeader + "." + tokenPayload
		signature := hmac.New(sha256.New, []byte(m.cnf.JwtSecret)).Sum([]byte(message))
		newSignature := util.Base64UrlEncoder(signature)

		if(tokenSignature != newSignature) {
			http.Error(w, "Unauthorized!!!", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}