package controlers

import (
	"Groupietracker/templates"
	"net/http"
)

func ConnectionControler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session-name")

	// Vérifier si l'utilisateur est déjà connecté
	if _, ok := session.Values["username"].(string); ok {
		// Rediriger vers la page profil
		http.Redirect(w, r, "/profil", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Exemple de vérification des identifiants (à remplacer par une vraie vérification)
		if username == "admin" && password == "admin" {
			// Enregistrer le nom d'utilisateur dans la session
			session.Values["username"] = username
			session.Save(r, w)

			// Rediriger vers la page profil
			http.Redirect(w, r, "/profil", http.StatusSeeOther)
			return
		}
	}

	// Afficher la page de connexion
	data := map[string]interface{}{
		"IsLoggedIn": false, // L'utilisateur n'est pas connecté
	}
	templates.ListTemp.ExecuteTemplate(w, "connection", data)
}
