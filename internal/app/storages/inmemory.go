package storages

import (
	"sync"
)

// InMemoryStore implements DataStore and holds URLs in memory
type InMemoryStore struct {
	sync.RWMutex
	URLs map[string]URLData
}

// NewInMemoryStore creates a new InMemoryStore
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{URLs: make(map[string]URLData)}
}

// AddURL adds a URL to the store
func (ims *InMemoryStore) AddURL(urlData URLData) error {
	ims.Lock()
	defer ims.Unlock()
	ims.URLs[urlData.ShortURL] = urlData
	return nil
}

// GetURL retrieves a URL by short URL
func (ims *InMemoryStore) GetURL(shortURL string) (URLData, bool) {
	ims.RLock()
	defer ims.RUnlock()
	urlData, exists := ims.URLs[shortURL]
	return urlData, exists
}

// DeleteURL deletes a URL from the store
func (ims *InMemoryStore) DeleteURL(shortURL string) error {
	ims.Lock()
	defer ims.Unlock()
	delete(ims.URLs, shortURL)
	return nil
}

// GetAllURLs returns all URLs
func (ims *InMemoryStore) GetAllURLs() []URLData {
	ims.RLock()
	defer ims.RUnlock()
	var urls []URLData
	for _, urlData := range ims.URLs {
		urls = append(urls, urlData)
	}
	return urls
}
