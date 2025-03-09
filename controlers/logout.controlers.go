package controlers

import (
	"net/http"
)

func LogoutControler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session-name")

	// Supprimer la session
	session.Options.MaxAge = -1 // Supprime le cookie de session
	session.Save(r, w)

	// Rediriger vers la page de connexion
	http.Redirect(w, r, "/connection", http.StatusSeeOther)
}
