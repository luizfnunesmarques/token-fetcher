

## Get rid of complexitt to fetch short lived tokens for your Github App installation

---
## Forget countless line long scripts
Provide the three env variables and good to go. Painless execution.

---
![Tests](https://github.com/luizfnunesmarques/token-fetcher/actions/workflows/lint-and-test.yml/badge.svg)
![Code report](https://goreportcard.com/badge/github.com/luizfnunesmarques-token-fetcher)
---
## Use cases
### Standalone binary
In need of testing against GitHub rest api as-is the application? Simply run the program in a terminal and vois-a-la, token's fetched.

### Automating a workflow
Need a installation token in any part of an automation proccess? The docker image is your friend. Refer to the examples folder for blueprints.

### I say I don't trust the image and secrets will be stolen :O
 Wise you. To guarantee the binary is generated after the actual code, compare the sha256 from the build phase with the actual binary embedded in the image.