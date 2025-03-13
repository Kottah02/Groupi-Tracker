package controlers

import (
	"Groupietracker/services"
	"Groupietracker/templates"
	"net/http"
	"strconv"
	"strings"
)

// créer une fonction controler qui permet de compléter la fonction route associé au meme template.
func CarteControler(w http.ResponseWriter, r *http.Request) {
	responsData, responseCode, responseErr := services.GetYugiohData()
	if responseErr != nil || responseCode != http.StatusOK {
		http.Redirect(w, r, "/error", http.StatusSeeOther) // Redirection au lieu de http.Error
		return
	}

	query := r.FormValue("query")
	attribute := r.FormValue("attribute")
	types := r.FormValue("type")
	LevelStr := r.FormValue("level")

	var levelInt int
	var err error
	if LevelStr != "" {
		levelInt, err = strconv.Atoi(LevelStr)
		if err != nil {
			http.Error(w, "Invalid level format", http.StatusBadRequest)
			return
		}
	}

	var cartesfiltre []services.YugiohCard
	for _, carte := range responsData.Data {
		filtreName := query == "" || strings.Contains(strings.ToLower(carte.Name), strings.ToLower(query))
		filtreAttribute := attribute == "" || attribute == carte.Attribute
		filtreType := types == "" || types == carte.Type
		filtreLevel := LevelStr == "" || carte.Level == levelInt

		if filtreName && filtreAttribute && filtreType && filtreLevel {
			cartesfiltre = append(cartesfiltre, carte)
		}
	}

	templates.ListTemp.ExecuteTemplate(w, "carte", cartesfiltre)
}
