package services

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, name string, value string, d time.Duration) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(d),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}
