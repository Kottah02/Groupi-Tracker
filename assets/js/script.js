// script.js

let currentIndex = 0;

// Fonction pour faire défiler le carrousel vers la gauche
function prevSlide() {
    const carouselInner = document.getElementById('carousel-inner');
    const totalItems = document.querySelectorAll('.carousel-item').length;
    currentIndex = (currentIndex - 1 + totalItems) % totalItems;
    carouselInner.style.transform = `translateX(-${currentIndex * 100}%)`;
}

// Fonction pour faire défiler le carrousel vers la droite
function nextSlide() {
    const carouselInner = document.getElementById('carousel-inner');
    const totalItems = document.querySelectorAll('.carousel-item').length;
    currentIndex = (currentIndex + 1) % totalItems;
    carouselInner.style.transform = `translateX(-${currentIndex * 100}%)`;
}

// Optionnel : Défilement automatique
function startAutoSlide() {
    setInterval(nextSlide, 1000); // Défile toutes les 5 secondes
}

// Démarrer le défilement automatique au chargement de la page
window.onload = startAutoSlide;