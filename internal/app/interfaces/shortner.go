package interfaces

import (
	"errors"
	"strings"

	"golang.org/x/exp/rand"
)

// URLShortener interface defines the methods for URL shortening
type URLShortener interface {
	Shorten(url string) (string, error)
	Retrieve(shortened string) (string, error)
}

// MapURLShortener is a struct that implements URLShortener interface
type MapURLShortener struct {
	urlMap       map[string]string
	shortenedMap map[string]string
}

// NewMapURLShortener creates a new instance of MapURLShortener
func NewMapURLShortener() *MapURLShortener {
	return &MapURLShortener{
		urlMap:       make(map[string]string),
		shortenedMap: make(map[string]string),
	}
}

// Shorten generates a shortened URL and stores it
func (m *MapURLShortener) Shorten(url string) (string, error) {
	if url == "" {
		return "", errors.New("URL cannot be empty")
	}

	shortened := GenerateShortURL() // Simple shortening logic
	m.urlMap[shortened] = url
	m.shortenedMap[url] = shortened
	return shortened, nil
}

// Retrieve gets the original URL for a shortened URL
func (m *MapURLShortener) Retrieve(shortened string) (string, error) {
	url, exists := m.urlMap[shortened]
	if !exists {
		return "", errors.New("shortened URL not found")
	}
	return url, nil
}

func GenerateShortURL() string {
	var charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var output strings.Builder
	length := 8
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
