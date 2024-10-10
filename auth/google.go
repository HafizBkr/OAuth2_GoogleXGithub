package auth

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "html/template"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

func HandleOAuthRedirect(w http.ResponseWriter, r *http.Request) {
    clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
    clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
    config := oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURL:  "http://localhost:8080/auth/callback",
        Scopes:       []string{"email", "profile"},
        Endpoint:     google.Endpoint,
    }
    url := config.AuthCodeURL("random", oauth2.AccessTypeOnline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleAuthCallback(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
    clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
    config := oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURL:  "http://localhost:8080/auth/callback",
        Scopes:       []string{"email", "profile"},
        Endpoint:     google.Endpoint,
    }
    token, err := config.Exchange(context.Background(), code)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    client := config.Client(context.Background(), token)
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des informations utilisateur", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var userInfo struct {
        Email          string `json:"email"`
        Name           string `json:"name"`
        ProfilePicture string `json:"picture"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        http.Error(w, "Erreur lors du décodage des informations utilisateur", http.StatusInternalServerError)
        return
    }
    if userInfo.Name == "" {
        http.Redirect(w, r, fmt.Sprintf("/complete-profile?email=%s", userInfo.Email), http.StatusSeeOther)
        return
    }
    fmt.Fprintf(w, "Bienvenue %s ! Profil complété.", userInfo.Name)
}

func HandleCompleteProfile(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")
    tmpl := template.Must(template.ParseFiles("static/complete-profile.html"))
    data := struct {
        Email string
    }{
        Email: email,
    }
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func SaveProfileHandler(w http.ResponseWriter, r *http.Request) {
    email := r.FormValue("email") 
    username := r.FormValue("username")
    file, _, err := r.FormFile("profile_picture")
    if err != nil {
        http.Error(w, "Erreur lors du téléchargement de la photo de profil", http.StatusBadRequest)
        return
    }
    defer file.Close()
    if username == "" {
        http.Error(w, "Le nom d'utilisateur est requis", http.StatusBadRequest)
        return
    }

    fmt.Printf("Nom d'utilisateur : %s, Email : %s\n", username, email) 
    http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}
