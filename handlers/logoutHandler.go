package handlers

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:        "Authorization",
		Value:       "",
		Quoted:      false,
		Path:        "/",
		Domain:      "",
		Expires:     time.Time{},
		RawExpires:  "",
		MaxAge:      0,
		Secure:      true,
		HttpOnly:    true,
		SameSite:    0,
		Partitioned: false,
		Raw:         "",
		Unparsed:    nil,
	})

	w.WriteHeader(http.StatusOK)
}
