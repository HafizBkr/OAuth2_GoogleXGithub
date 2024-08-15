package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)


func HandleOAuthRedirectGit(w http.ResponseWriter, r *http.Request) {
       clientID := os.Getenv("GITHUB_OAUTH_CLIENT_ID")
       clientSecret := os.Getenv("GITHUB_OAUTH_CLIENT_SECRET")
       config := oauth2.Config{
           ClientID:     clientID,
           ClientSecret: clientSecret,
           RedirectURL:  "http://localhost:8080/auth/github/callback",
           Scopes:       []string{"email", "profile"},
           Endpoint:     github.Endpoint,
       }
       url := config.AuthCodeURL("random", oauth2.AccessTypeOnline)
       http.Redirect(w, r, url, http.StatusTemporaryRedirect)
   }
   
   func HandleAuthCallbackGit(w http.ResponseWriter, r *http.Request) {
       code := r.URL.Query().Get("code")
       clientID := os.Getenv("GITHUB_OAUTH_CLIENT_ID")
       clientSecret := os.Getenv("GITHUB_OAUTH_CLIENT_SECRET")
       config := oauth2.Config{
           ClientID:     clientID,
           ClientSecret: clientSecret,
           RedirectURL:  "http://localhost:8080/auth/github/callback",
           Scopes:       []string{"email", "profile"},
           Endpoint:     github.Endpoint,
       }
       token, err := config.Exchange(context.Background(), code)
       if err != nil {
           http.Error(w, err.Error(), http.StatusInternalServerError)
           return
       }
       fmt.Printf("Token obtained: %s\n", token.AccessToken)
       return
   }