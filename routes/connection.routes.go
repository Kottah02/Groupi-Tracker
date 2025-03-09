package routes

import (
	"Groupietracker/controlers"
	"net/http"
)

func SetupConnectionRoutes() {
	http.HandleFunc("/connection", controlers.ConnectionControler)
	http.HandleFunc("/profil", controlers.ProfilControler)
	http.HandleFunc("/logout", controlers.LogoutControler) // Ajouter la route de déconnexion
}
