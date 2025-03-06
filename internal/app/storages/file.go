package storages

import (
	"encoding/json"
	"os"
)

// FileStore implements DataStore and holds URLs in a file
type FileStore struct {
	FilePath string
	URLs     map[string]URLData
}

// NewFileStore creates a new FileStore
func NewFileStore(filePath string) *FileStore {
	return &FileStore{FilePath: filePath, URLs: make(map[string]URLData)}
}

// LoadData loads URL data from the file, creating an empty file if it does not exist
func (fs *FileStore) LoadData() error {
	if _, err := os.Stat(fs.FilePath); os.IsNotExist(err) {
		// Create an empty file
		emptyData := make(map[string]URLData)
		data, _ := json.Marshal(emptyData)
		if err := os.WriteFile(fs.FilePath, data, 0644); err != nil {
			return err
		}
		fs.URLs = emptyData // Initialize URLs to empty
		return nil
	}

	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &fs.URLs)
}

// SaveData saves URL data to the file
func (fs *FileStore) SaveData() error {
	data, err := json.Marshal(fs.URLs)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, data, 0644)
}

// AddURL adds a URL to the store
func (fs *FileStore) AddURL(urlData URLData) error {
	fs.URLs[urlData.ShortURL] = urlData
	return fs.SaveData()
}

// GetURL retrieves a URL by short URL
func (fs *FileStore) GetURL(shortURL string) (URLData, bool) {
	urlData, exists := fs.URLs[shortURL]
	return urlData, exists
}

// DeleteURL deletes a URL from the store
func (fs *FileStore) DeleteURL(shortURL string) error {
	delete(fs.URLs, shortURL)
	return fs.SaveData()
}

// GetAllURLs returns all URLs
func (fs *FileStore) GetAllURLs() []URLData {
	var urls []URLData
	for _, urlData := range fs.URLs {
		urls = append(urls, urlData)
	}
	return urls
}
