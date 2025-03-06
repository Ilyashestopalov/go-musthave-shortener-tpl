package handlers

import (
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/storages"
)

// URLHandler handles URL shortening and retrieval
type URLHandler struct {
	store   storages.DataStore
	baseURL string
}

// NewURLHandler creates a new URLHandler
func NewURLHandler(store storages.DataStore, baseURL string) *URLHandler {
	return &URLHandler{store: store, baseURL: baseURL}
}
