package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey returns the API key from the request headers
// Authorization : ApiKey <key>
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No Authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed Authorization first part of header")
	}
	return vals[1], nil
}
