package auth

import (
	// Go Builtin Packages
	"log"
	"net/http"

	// Google Appengine Packages
	"appengine"
	"appengine/user"
)

func Login(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	dest := r.Referer()

	if usr := user.Current(c); usr != nil { // already logged in
		if dest != "/pastebin/login/" {
			http.Redirect(w, r, dest, http.StatusFound)
		} else {
			http.Redirect(w, r, "/pastebin/", http.StatusFound)
		}
	} else {
		if url, err := user.LoginURL(c, dest); err == nil {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			log.Panicln(err)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	dest := r.Referer()

	if usr := user.Current(c); usr == nil { // already logged out
		if dest != "/pastebin/logout/" {
			http.Redirect(w, r, dest, http.StatusFound)
		} else {
			http.Redirect(w, r, "/pastebin/", http.StatusFound)
		}
	} else {
		if url, err := user.LogoutURL(c, dest); err == nil {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			log.Panicln(err)
		}
	}
}