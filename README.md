## get rid of complexity to fetch short lived tokens for your Github App installation
Token fetcher generates
---
## forget countless lines long scripts
Token fetcher only requires three env variables set on the executing shell. It handles gracefully most common issues when getting the tokens. No need to jq the response for the request made with curl.

---
![Build and release](https://img.shields.io/github/v/release/luizfnunesmarques/token-fetcher.svg)
![Tests](https://github.com/luizfnunesmarques/token-fetcher/actions/workflows/lint-and-test.yml/badge.svg)
![Code report](https://goreportcard.com/badge/github.com/luizfnunesmarques-token-fetcher)

---
## use cases

### standalone binary
In need of testing against GitHub rest api as-is the application? Simply run the program in a terminal and voilà, token's fetched.

### docker image
Need a installation token in any part of an automation proccess? The docker image is your friend. Refer to the examples folder for blueprints.

---
## Usage/running

Token fetcher reads three environment variables and use them to authenticate and sign the request to GitHub's API: APP_ID, INSTALLATION_ID and PRIVATE_KEY.

### From source
Build the project:
```
go build .
```

```
export INSTALLATION_ID= APP_ID= PRIVATE_KEY=
```

ps: Note the project uses go modules and so the GO111MODULE variable should be `on`.

### From the docker image
`docker pull luizfnunesmarques:token-fetcher`

---
## useful links:
Authenticating as a GitHub installation: https://docs.github.com/en/developers/apps/building-github-apps/authenticating-with-github-apps

GO111MODULE variable: https://maelvls.dev
