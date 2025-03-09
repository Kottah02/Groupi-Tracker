document.getElementById('loginForm').addEventListener('submit', function (event) {
    event.preventDefault(); // Empêcher la soumission du formulaire

    // Récupérer les valeurs des champs
    const username = document.getElementById('username').value.trim();
    const password = document.getElementById('password').value.trim();

    // Vérifier les identifiants (exemple simple)
    if (username === "admin" && password === "password") {
        // Rediriger vers la page d'accueil ou une autre page après connexion réussie
        window.location.href = "/";
    } else {
        // Afficher un message d'erreur
        document.getElementById('errorMessage').textContent = "Nom d'utilisateur ou mot de passe incorrect.";
    }
});