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

func getGithubClientID() string {
    githubClientID, exists := os.LookupEnv("GITHUB_CLIENT_ID")
    if !exists {
        log.Fatal("Github Client ID not defined in .env file")
    }
    return githubClientID
}

func getGithubClientSecret() string {
    githubClientSecret, exists := os.LookupEnv("GITHUB_CLIENT_SECRET")
    if !exists {
        log.Fatal("Github Client Secret not defined in .env file")
    }
    return githubClientSecret
}

func getGithubAccessToken(code string) string {
    clientID := getGithubClientID()
    clientSecret := getGithubClientSecret()

    requestBody := map[string]string{
        "client_id":     clientID,
        "client_secret": clientSecret,
        "code":          code,
    }
    requestJSON, _ := json.Marshal(requestBody)

    req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON))
    if err != nil {
        log.Panic("Request creation failed")
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Panic("Request failed")
    }
    defer resp.Body.Close()

    respBody, _ := ioutil.ReadAll(resp.Body)

    type githubAccessTokenResponse struct {
        AccessToken string `json:"access_token"`
    }

    var ghResp githubAccessTokenResponse
    json.Unmarshal(respBody, &ghResp)

    return ghResp.AccessToken
}

func getGithubData(accessToken string) string {
    req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
    if err != nil {
        log.Panic("API Request creation failed")
    }

    req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Panic("Request failed")
    }
    defer resp.Body.Close()

    respBody, _ := ioutil.ReadAll(resp.Body)
    return string(respBody)
}
