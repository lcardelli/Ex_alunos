package auth

import (
    "net/http"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var (
    googleOauthConfig = &oauth2.Config{
        RedirectURL:  "http://localhost:8080/callback",
        ClientID:     "YOUR_CLIENT_ID",
        ClientSecret: "YOUR_CLIENT_SECRET",
        Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint:     google.Endpoint,
    }
)

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
    url := googleOauthConfig.AuthCodeURL("random-state")
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
    // Processar o token recebido e verificar a autenticação
}
