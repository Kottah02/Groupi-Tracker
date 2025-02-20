package controlers

import (
	"Groupietracker/templates"
	"net/http"
)

// créer une fonction controler qui permet de compléter la fonction route associé au meme template.
func ArticleControler(w http.ResponseWriter, r *http.Request) {
	templates.ListTemp.ExecuteTemplate(w, "article", nil)
}
