package controlers

import (
	"Groupietracker/templates"
	"net/http"
)

func ProfilControler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session-name")

	// Vérifier si l'utilisateur est connecté
	if username, ok := session.Values["username"].(string); ok {
		// Afficher la page profil avec le nom d'utilisateur
		data := map[string]interface{}{
			"Username":   username,
			"IsLoggedIn": true, // L'utilisateur est connecté
		}
		templates.ListTemp.ExecuteTemplate(w, "profil", data)
		return
	}

	// Rediriger vers la page de connexion si l'utilisateur n'est pas connecté
	http.Redirect(w, r, "/connection", http.StatusSeeOther)
}
