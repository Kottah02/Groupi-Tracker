package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	// Servir les fichiers statiques (CSS, images, etc.) depuis le dossier "assets"
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Définir les routes pour les différentes pages
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/classe", serveClasse)
	http.HandleFunc("/monture", serveMonture)

	// Démarrer le serveur sur le port 8080
	fmt.Println("Le serveur est lancé sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur : ", err)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	// Servir le fichier index.html depuis le dossier "template"
	http.ServeFile(w, r, filepath.Join("template", "index.html"))
}

func serveClasse(w http.ResponseWriter, r *http.Request) {
	// Servir le fichier classe.html depuis le dossier "template"
	http.ServeFile(w, r, filepath.Join("template", "classe.html"))
}

func serveMonture(w http.ResponseWriter, r *http.Request) {
	// Servir le fichier monture.html depuis le dossier "template"
	http.ServeFile(w, r, filepath.Join("template", "monture.html"))
}
