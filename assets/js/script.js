let allCards = []; // Stocke toutes les cartes récupérées
let currentPage = 0; // Page actuelle
const cardsPerPage = 20; // Nombre de cartes par page

// Fonction pour charger les cartes en fonction du fragment de nom, de l'attribut, du type et du niveau
function loadCards(searchTerm, attribute, type, level) {
    let url = 'https://db.ygoprodeck.com/api/v7/cardinfo.php';

    // Si un terme de recherche est fourni, l'ajouter à l'URL
    if (searchTerm) {
        url += `?fname=${encodeURIComponent(searchTerm)}`;
    }

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Erreur HTTP: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            allCards = data.data || []; // Stocker toutes les cartes

            // Filtrer les cartes par attribut si un attribut est sélectionné
            if (attribute) {
                allCards = allCards.filter(card => card.attribute === attribute);
            }

            // Filtrer les cartes par type si un type est sélectionné
            if (type) {
                allCards = allCards.filter(card => card.type === type);
            }

            // Filtrer les cartes par niveau si un niveau est sélectionné
            if (level) {
                allCards = allCards.filter(card => card.level == level); // Utilisez == pour la comparaison (niveau est un nombre)
            }

            currentPage = 0; // Réinitialiser la page à 0
            displayCards(); // Afficher les cartes de la première page
            updatePageInfo(); // Mettre à jour l'affichage du numéro de page
        })
        .catch(error => {
            console.error('Erreur:', error);
            document.getElementById('results').innerHTML = `<p>Erreur lors de la recherche : ${error.message}</p>`;
        });
}

// Fonction pour afficher les cartes de la page actuelle
function displayCards() {
    const resultsDiv = document.getElementById('results');
    resultsDiv.innerHTML = ''; // Effacer les résultats précédents

    // Calculer les indices de début et de fin pour la page actuelle
    const startIndex = currentPage * cardsPerPage;
    const endIndex = startIndex + cardsPerPage;
    const cardsToDisplay = allCards.slice(startIndex, endIndex);

    // Afficher les cartes de la page actuelle
    if (cardsToDisplay.length > 0) {
        cardsToDisplay.forEach(card => {
            const cardElement = document.createElement('div');
            cardElement.classList.add('card-result');

            // Afficher l'image, le nom, le type, l'attribut, les statistiques et la description
            cardElement.innerHTML = `
                <img src="${card.card_images[0].image_url_small}" alt="${card.name}">
                <div class="card-info">
                    <h3>${card.name}</h3>
                    <p class="type"><strong>Type :</strong> ${card.type}</p>
                    ${card.attribute ? `<p class="attribute"><strong>Attribut :</strong> ${card.attribute}</p>` : ''}
                    ${card.level ? `<p class="level"><strong>Niveau :</strong> ${card.level}</p>` : ''}
                    <div class="stats">
                        ${card.atk ? `<p><strong>ATK :</strong> ${card.atk}</p>` : ''}
                        ${card.def ? `<p><strong>DEF :</strong> ${card.def}</p>` : ''}
                    </div>
                    <p class="description"><strong>Description :</strong> ${card.desc}</p>
                </div>
            `;

            resultsDiv.appendChild(cardElement);
        });
    } else {
        resultsDiv.innerHTML = '<p>Aucune carte trouvée.</p>';
    }

    // Mettre à jour l'état des boutons de pagination
    document.getElementById('prevPage').disabled = currentPage === 0;
    document.getElementById('nextPage').disabled = endIndex >= allCards.length;

    // Mettre à jour l'affichage du numéro de page
    updatePageInfo();
}

// Fonction pour mettre à jour l'affichage du numéro de page
function updatePageInfo() {
    const totalPages = Math.ceil(allCards.length / cardsPerPage);
    const pageInfo = document.getElementById('pageInfo');
    pageInfo.textContent = `${currentPage + 1}/${totalPages}`;
}

// Fonction pour gérer la recherche
function searchCard() {
    const searchTerm = document.getElementById('searchBar').value.trim();
    const attribute = document.getElementById('attributeFilter').value;
    const type = document.getElementById('typeFilter').value;
    const level = document.getElementById('levelFilter').value;

    loadCards(searchTerm, attribute, type, level);
}

// Gestion des boutons de pagination
document.getElementById('prevPage').addEventListener('click', () => {
    if (currentPage > 0) {
        currentPage--;
        displayCards();
    }
});

document.getElementById('nextPage').addEventListener('click', () => {
    const totalPages = Math.ceil(allCards.length / cardsPerPage);
    if (currentPage < totalPages - 1) {
        currentPage++;
        displayCards();
    }
});

// Gestion du bouton de recherche
document.getElementById('searchButton').addEventListener('click', searchCard);