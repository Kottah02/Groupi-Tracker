// Types de monstres valides
const monsterTypes = [
    'Normal Monster',
    'Effect Monster',
    'Tuner Monster',
    'Flip Monster',
    'Spirit Monster',
    'Union Effect Monster',
    'Gemini Monster',
    'Pendulum Effect Monster',
    'Ritual Monster',
    'Fusion Monster',
    'Synchro Monster',
    'XYZ Monster',
    'Link Monster'
  ];
  
  // Récupérer les données pour chaque type de monstre
  Promise.all(
    monsterTypes.map(type =>
      fetch(`https://db.ygoprodeck.com/api/v7/cardinfo.php?type=${type}`)
        .then(response => response.json())
        .then(data => data.data)
    )
  )
    .then(results => {
      // Fusionner tous les résultats en un seul tableau
      const allMonsters = results.flat();
  
      // Trier les monstres par niveau (ordre croissant)
      const sortedMonsters = allMonsters.sort((a, b) => a.level - b.level);
  
      // Afficher les résultats
      console.log(sortedMonsters);
    })
    .catch(error => console.error('Erreur lors de la récupération des données :', error));