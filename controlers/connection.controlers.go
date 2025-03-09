package controlers

import (
	"Groupietracker/templates"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func ConnectionControler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	// Vérifier si l'utilisateur est déjà connecté
	if username, ok := session.Values["username"].(string); ok {
		data := map[string]string{
			"Username": username,
		}
		templates.ListTemp.ExecuteTemplate(w, "connection", data)
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

			// Rediriger pour éviter la resoumission du formulaire
			http.Redirect(w, r, "/connection", http.StatusSeeOther)
			return
		}
	}

	// Afficher la page de connexion
	templates.ListTemp.ExecuteTemplate(w, "connection", nil)
}
