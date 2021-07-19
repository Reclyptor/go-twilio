# go-twilio
Twilio API Client \
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/Reclyptor/go-twilio?color=blue&label=Release&sort=semver&style=plastic)
![GitHub](https://img.shields.io/github/license/Reclyptor/go-twilio?color=red&label=License&style=plastic)
![GitHub repo size](https://img.shields.io/github/repo-size/Reclyptor/go-twilio?color=green&label=Size&style=plastic)

## Installation
```shell script
go get -d -u github.com/Reclyptor/go-twilio
```

## Sample Usage
```go
package main

import (
	"github.com/reclyptor/go-twilio"
	"os"
)

var ACCOUNT_SID = os.Getenv("ACCOUNT_SID")
var AUTH_TOKEN = os.Getenv("AUTH_TOKEN")
var PHONE_NUMBER = os.Getenv("PHONE_NUMBER")

func main() {
	client := twilio.NewClient(ACCOUNT_SID, AUTH_TOKEN, PHONE_NUMBER)
	client.SendSMS("<< PHONE_NUMBER >>", "<< MESSAGE >>")
}
```
