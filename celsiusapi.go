package main

// I'm basically just copying this guy's kraken code to learn Go
// https://github.com/beldur/kraken-go-api-client/blob/master/krakenapi.go

import (
	// "crypto/hmac"
	// "crypto/sha256"
	// "crypto/sha512"
	// "encoding/base64"
	// "encoding/json"
	// "errors"
	"fmt"
	// "io/ioutil"
	// "math/big"
	// "mime"
	"net/http"
	// "net/url"
	// "strconv"
	// "strings"
	// "time"
)


const (
	// StagingBaseURL is the official Celsius API endpoint
	StagingBaseURL = "https://wallet-api.staging.celsius.network"
	// APIVersion is the official Celsius API version
	APIVersion = "0.1.0"
	// Author because this is my secret kingly throne
	Author = "Dennis Dang"
)

// CelsiusAPI is the Celsius API connection
type CelsiusAPI struct {
	key string
	secret string
	client *http.Client
}

// New creates the Celsius API client
func New(key string, secret string) *CelsiusAPI {
	return &CelsiusAPI{key, secret, http.DefaultClient}
}

func main() {
	fmt.Println("Celsius client is ready to go!");
}

// HEALTH
// WALLET
// KYC
// UTIL
// INSTITUTIONAL
