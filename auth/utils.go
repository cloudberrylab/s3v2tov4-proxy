package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/url"
	"strings"
)

// sum256 calculate sha256 sum for an input byte array.
func sum256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

// sumHMAC calculate hmac between two input byte array.
func sumHMAC(key []byte, data []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

func sortQuery(encodedQuery string) string {
	m, _ := url.ParseQuery(encodedQuery)
	return strings.Replace(m.Encode(), "+", "%20", -1)
}
