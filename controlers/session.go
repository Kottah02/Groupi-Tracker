package controlers

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("secret-key"))

func init() {
	// Configurer la session pour qu'elle expire à la fermeture du navigateur
	Store.Options = &sessions.Options{
		MaxAge:   0,    // Session expire à la fermeture du navigateur
		HttpOnly: true, // Empêche l'accès au cookie via JavaScript
	}
}
