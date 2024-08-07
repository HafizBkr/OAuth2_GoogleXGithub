package auth

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "bytes"
)

func getGoogleClientID() string {
    googleClientID, exists := os.LookupEnv("GOOGLE_CLIENT_ID")
    if !exists {
        log.Fatal("Google Client ID not defined in .env file")
    }
    return googleClientID
}

func getGoogleClientSecret() string {
    googleClientSecret, exists := os.LookupEnv("GOOGLE_CLIENT_SECRET")
    if !exists {
        log.Fatal("Google Client Secret not defined in .env file")
    }
    return googleClientSecret
}

func getGoogleAccessToken(code string) string {
    clientID := getGoogleClientID()
    clientSecret := getGoogleClientSecret()

    requestBody := map[string]string{
        "client_id":     clientID,
        "client_secret": clientSecret,
        "code":          code,
        "grant_type":    "authorization_code",
        "redirect_uri":  "http://localhost:8080/auth/google/callback",
    }
    requestJSON, _ := json.Marshal(requestBody)

    req, err := http.NewRequest("POST", "https://oauth2.googleapis.com/token", bytes.NewBuffer(requestJSON))
    if err != nil {
        log.Panic("Request creation failed")
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Panic("Request failed")
    }
    defer resp.Body.Close()

    respBody, _ := ioutil.ReadAll(resp.Body)

    type googleAccessTokenResponse struct {
        AccessToken string `json:"access_token"`
    }

    var googleResp googleAccessTokenResponse
    json.Unmarshal(respBody, &googleResp)

    return googleResp.AccessToken
}

func getGoogleData(accessToken string) string {
    req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
    if err != nil {
        log.Panic("API Request creation failed")
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Panic("Request failed")
    }
    defer resp.Body.Close()

    respBody, _ := ioutil.ReadAll(resp.Body)
    return string(respBody)
}
