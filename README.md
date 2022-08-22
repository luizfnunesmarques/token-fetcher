## get rid of complexity to fetch short lived tokens 🌸

Fetching an installation token can be useful and quite stressful. Token fetcher lowers the complexity: it only requires three env variables set on the executing shell.
Handling gracefully most common issues as unauthorised, unauthenticated or "network has issues", it only has one output: the token.

---
![Build and release](https://img.shields.io/github/v/release/luizfnunesmarques/token-fetcher.svg)
![Tests](https://github.com/luizfnunesmarques/token-fetcher/actions/workflows/lint-and-test.yml/badge.svg)
![Code report](https://goreportcard.com/badge/github.com/luizfnunesmarques-token-fetcher)
![Docker pulls](https://img.shields.io/docker/pulls/luizfnunesmarques/token-fetcher.svg)
---
## use cases
### Daily operations for devs
In need of testing against GitHub rest api as-is or want to check some metadata returned from it? Shell, `./token-fetcher` and token is ready. Unlimited ~power~ simplicity.

### Automation
In need to push to a repository, manipulate metadata or really any other thing as the Github application? The docker image is your friend. Refer to the examples folder on common scenarios.

---
## usage & running

Token fetcher reads three environment variables and use them to authenticate and sign the request to GitHub's API: APP_ID, INSTALLATION_ID and PRIVATE_KEY.
#### From source
Build the project:
`go build .`

Define the variables and run it:
```
export INSTALLATION_ID= APP_ID= PRIVATE_KEY=
./token-fetcher
```
ps: Note the project uses go modules and so the GO111MODULE variable should be `on`.
#### From the docker image locally (or building your own)
`docker run -e APP_ID= -e INSTALLATION_ID= -e PRIVATE_KEY= luizfnunesmarques:token-fetcher`
