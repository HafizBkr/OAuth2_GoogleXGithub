#  Oauth2 Google and Github
markdown
Copier le code
# OAuth2_GoogleXGithub

Projet d'authentification OAuth2 intégrant les connexions avec Google et GitHub, développé en Go.

## Structure du Projet


        .
        ├── auth
        │   ├── github.go        # Gestion de l'authentification avec GitHub
        │   └── google.go        # Gestion de l'authentification avec Google
        ├── env_exemple          # Exemple de fichier d'environnement (.env)
        ├── go.mod               # Fichier de dépendances Go
        ├── go.sum               # Somme de contrôle pour les dépendances Go
        ├── main.go              # Point d'entrée de l'application
        ├── providers
        │   ├── email.go         # Gestion des emails
        │   └── jwt.go           # Gestion des tokens JWT
        ├── README.md            # Documentation du projet
        ├── static
        │   └── index.html       # Page HTML statique pour les tests
        └── tmp
            ├── build-errors.log # Fichier de log pour les erreurs de build
            └── main             # Binaire généré (exemple)
            
#   Installation
Clonez le dépôt :

          git clone https://github.com/HafizBkr/OAuth2_GoogleXGithub.git
Accédez au répertoire du projet :

            cd OAuth2_GoogleXGithub
Installez les dépendances :


          go mod download
          
Configuration:
Créez un fichier .env en suivant l'exemple fourni dans env_exemple.
Ajoutez vos identifiants OAuth2 pour Google et GitHub dans le fichier .env (évitez de les committer pour des raisons de sécurité).
Exemple .env :


      GOOGLE_CLIENT_ID=your_google_client_id
      GOOGLE_CLIENT_SECRET=your_google_client_secret
      GITHUB_CLIENT_ID=your_github_client_id
      GITHUB_CLIENT_SECRET=your_github_client_secret
      JWT_SECRET=your_jwt_secret
Utilisation
Lancez l'application :

    go run main.go
    
Ouvrez un navigateur et accédez à `http://localhost:8080/static/index.html pour tester` les connexions OAuth2.

Fonctionnalités
Auth Google : Authentification via Google.
Auth GitHub : Authentification via GitHub.
JWT Tokens : Génération et gestion des tokens JWT.

Gestion des emails : Fournisseur d'email pour les notifications (non configuré par défaut).

#  Contribuer
Forkez le projet.
Créez une branche pour votre fonctionnalité (`git checkout -b feature/new-feature`).
Commitez vos modifications (`git commit -am 'Add new feature'`).
Poussez la branche (`git push origin feature/new-feature`).
Ouvrez une pull request.
