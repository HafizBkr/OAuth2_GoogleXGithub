package auth

import(
    "fmt"
	"os"
	"net/http"
    "context"
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
       fmt.Printf("Token obtained: %s\n", token.AccessToken)
       return
   }