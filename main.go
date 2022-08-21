package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/go-github/v45/github"
)

type installationTokenResponse struct {
	Token        string                         `json:"token"`
	ExpiresAt    time.Time                      `json:"expires_at"`
	Permissions  github.InstallationPermissions `json:"permissions,omitempty"`
	Repositories []github.Repository            `json:"repositories,omitempty"`
}

func fetchInstallationToken(installationID, appID string, privateKey []byte) (string, error) {
	claims := &jwt.StandardClaims{
		IssuedAt:  time.Now().Unix() - 60,
		ExpiresAt: time.Now().Unix() + 120,
		Issuer:    appID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	if err != nil {
		return "", err
	}

	ss, err := token.SignedString(signKey)

	if err != nil {
		return "", err
	}

	bearer := fmt.Sprintf("Bearer %s", ss)

	accessTokenEndpoint := fmt.Sprintf("https://api.github.com/app/installations/%s/access_tokens", installationID)

	req, _ := http.NewRequest(http.MethodPost, accessTokenEndpoint, nil)

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if response.StatusCode != 201 {
		return "", errors.New(response.Status)
	}

	defer response.Body.Close()
	var tokenResponse installationTokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)

	if err != nil {
		return "", err
	}

	return tokenResponse.Token, nil
}

func main() {
	getEnv := func(key string) string {
		val, envExists := os.LookupEnv(key)

		if !envExists {
			log.Fatalf("Missing env: %s", key)
		}

		return val
	}

	installationID := getEnv("INSTALLATION_ID")
	appID := getEnv("APP_ID")
	privateKey := getEnv("PRIVATE_KEY")

	token, err := fetchInstallationToken(installationID, appID, []byte(privateKey))

	if err != nil {
		log.Fatalf("Fetch failed with: %s", err)
	}

	fmt.Print(token)
}
