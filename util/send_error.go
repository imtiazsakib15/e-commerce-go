package util

import (
	"net/http"
)

func SendError(w http.ResponseWriter, statusCode int, err interface{}) {
	http.Error(w, err.(error).Error(), http.StatusBadRequest)
}