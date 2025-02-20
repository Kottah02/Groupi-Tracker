package routes

import (
	"Groupietracker/controlers"
	"net/http"
)

// Créer une route home qui va serveur de page d'acceuil, le "/" et utiliser uniquement pour la page d'acceuil sinon changer par le nom du template.
func carteRouter() {
	http.HandleFunc("/carte", controlers.CarteControler)
}
