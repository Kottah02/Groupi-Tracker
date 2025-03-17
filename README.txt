Documentation du Projet Groupie Tracker avec l'API Yu-Gi-Oh!

1. Introduction
Groupie Tracker est une application web qui permet aux utilisateurs de rechercher et de visualiser des informations sur les cartes Yu-Gi-Oh!
en utilisant l'API publique de Yu-Gi-Oh! (YGOPRODeck). 
L'application offre une interface intuitive pour explorer les cartes, 
filtrer les résultats, et afficher les détails de chaque carte, y compris les images, les attributs, les types, et les niveaux.

2. Fonctionnalités
2.1 Fonctionnalités principales
Recherche de cartes : Les utilisateurs peuvent rechercher des cartes par nom.

Filtres : Les utilisateurs peuvent filtrer les cartes par :

Attribut (par exemple, DARK, LIGHT, EARTH, etc.).

Type (par exemple, Monster, Spell, Trap).

Niveau (par exemple, 1 à 12).

Affichage des détails : Pour chaque carte, les utilisateurs peuvent voir :

Le nom de la carte.

Une image de la carte.

L'attribut, le type, le niveau, l'ATK, et la DEF.

La description de la carte.

Pagination : Les résultats de recherche sont paginés pour une meilleure expérience utilisateur.

2.2 Fonctionnalités supplémentaires
Interface utilisateur intuitive : Une interface utilisateur conviviale et responsive.

Mode sombre : Option pour basculer entre un mode clair et un mode sombre.

Connexion utilisateur (optionnelle) : Les utilisateurs peuvent se connecter pour enregistrer leurs cartes préférées.

3.1 Technologies utilisées
Backend :

Go (Golang) : Langage de programmation pour le serveur.

API Yu-Gi-Oh! : Utilisée pour récupérer les données des cartes.

Frontend :

HTML/CSS : Pour la structure et le style de l'interface utilisateur.

4. Installation et Exécution
4.1 Prérequis
Go : Installé sur votre machine. Vous pouvez le télécharger depuis golang.org.

Git : Pour cloner le dépôt du projet.

4.2 Étapes d'installation
Cloner le dépôt :

bash
Copy
git clone https://github.com/Kottah02/Groupi-Tracker.git
cd groupie-tracker
Installer les dépendances :

Aucune dépendance externe n'est requise pour Go, sauf si vous utilisez une base de données.

Configurer l'API :

Assurez-vous que l'API Yu-Gi-Oh! est accessible et que les endpoints sont correctement configurés dans le fichier handlers.go.

Lancer le serveur :

bash
Copy
go run main.go
Accéder à l'application :

Ouvrez votre navigateur et accédez à http://localhost:8080.

5. Utilisation de l'Application
5.1 Page d'accueil
La page d'accueil affiche une liste de cartes populaires.

Les utilisateurs peuvent cliquer sur une carte pour voir plus de détails.

5.2 Recherche et Filtres
Utilisez la barre de recherche pour trouver une carte par nom.

Appliquez des filtres pour affiner les résultats (par attribut, type, niveau, etc.).

5.3 Détails de la carte
Cliquez sur une carte pour voir :

Le nom de la carte.

Une image de la carte.

L'attribut, le type, le niveau, l'ATK, et la DEF.

La description de la carte.

6. Documentation de l'API Yu-Gi-Oh!
6.1 Endpoints utilisés
Récupérer toutes les cartes :

Copy
GET https://db.ygoprodeck.com/api/v7/cardinfo.php
Filtrer les cartes par nom :

Copy
GET https://db.ygoprodeck.com/api/v7/cardinfo.php?name={nom de la carte}
Filtrer les cartes par attribut :

Copy
GET https://db.ygoprodeck.com/api/v7/cardinfo.php?attribute={attribut}
Filtrer les cartes par type :

Copy
GET https://db.ygoprodeck.com/api/v7/cardinfo.php?type={type}
6.2 Exemple de réponse JSON
json
Copy

{
  "data": [
    {
      "id": 46986414,
      "name": "Dark Magician",
      "type": "Spellcaster",
      "level": 7,
      "atk": 2500,
      "def": 2100,
      "race": "Spellcaster",
      "attribute": "DARK",
      "card_images": [
        {
          "image_url": "https://storage.googleapis.com/ygoprodeck.com/pics/46986414.jpg",
          "image_url_small": "https://storage.googleapis.com/ygoprodeck.com/pics_small/46986414.jpg"
        }
      ]
    }
  ]
}

7. Améliorations futures
Connexion utilisateur : Permettre aux utilisateurs de se connecter et d'enregistrer leurs cartes préférées.

Mode sombre : Ajouter un mode sombre pour améliorer l'expérience utilisateur.

Optimisation des performances : Améliorer les temps de chargement et l'efficacité des requêtes API.

8. Conclusion
Groupie Tracker est une application web puissante et intuitive pour explorer les cartes Yu-Gi-Oh!.
Grâce à son interface utilisateur conviviale et ses fonctionnalités de recherche et de filtrage avancées,
les utilisateurs peuvent facilement trouver et visualiser les informations dont ils ont besoin.
Ce projet démontre l'utilisation efficace de Go pour le backend et des technologies modernes pour le frontend.