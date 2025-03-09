// Fonction pour rechercher une carte par son nom
function searchCard() {
  const searchTerm = document.getElementById('searchBar').value.trim();

  // Vérifier si le terme de recherche est vide
  if (searchTerm === "") {
      document.getElementById('results').innerHTML = '<p>Veuillez entrer un nom de carte.</p>';
      return;
  }

  // Construire l'URL de l'API avec le terme de recherche
  const url = `https://db.ygoprodeck.com/api/v7/cardinfo.php?fname=${encodeURIComponent(searchTerm)}`;

  // Envoyer la requête à l'API
  fetch(url)
      .then(response => {
          if (!response.ok) {
              throw new Error(`Erreur HTTP: ${response.status}`);
          }
          return response.json();
      })
      .then(data => {
          const resultsDiv = document.getElementById('results');
          resultsDiv.innerHTML = ''; // Effacer les résultats précédents

          // Vérifier si des cartes ont été trouvées
          if (data.data && data.data.length > 0) {
              data.data.forEach(card => {
                  const cardElement = document.createElement('div');
                  cardElement.classList.add('card-result');

                  // Afficher le nom, l'image et la description de la carte
                  cardElement.innerHTML = `
                      <img src="${card.card_images[0].image_url_small}" alt="${card.name}">
                      <div class="card-info">
                          <h3>${card.name}</h3>
                          <p class="type">${card.type}</p>
                          <div class="stats">
                              <p>ATK: ${card.atk || 'N/A'}</p>
                              <p>DEF: ${card.def || 'N/A'}</p>
                          </div>
                      </div>
                  `;

                  resultsDiv.appendChild(cardElement);
              });
          } else {
              resultsDiv.innerHTML = '<p>Aucune carte trouvée.</p>';
          }
      })
      .catch(error => {
          console.error('Erreur:', error);
          document.getElementById('results').innerHTML = `<p>Erreur lors de la recherche : ${error.message}</p>`;
      });
}

// Ajouter un écouteur d'événement pour le bouton de recherche
document.getElementById('searchButton').addEventListener('click', searchCard);