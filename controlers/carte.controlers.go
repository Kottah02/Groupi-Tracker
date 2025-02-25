package controlers

import (
	"Groupietracker/services"
	"Groupietracker/templates"
	"net/http"
)

// créer une fonction controler qui permet de compléter la fonction route associé au meme template.
func CarteControler(w http.ResponseWriter, r *http.Request) {
	responsData, responseCode, responseErr := services.GetYugiohData()
	if responseErr != nil || responseCode != http.StatusOK {
		http.Error(w, responseErr.Error(), http.StatusInternalServerError) // a changer par une redirection !
		return
	}

	templates.ListTemp.ExecuteTemplate(w, "carte", responsData)
}
