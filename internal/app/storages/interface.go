package storages

// DataStore is an interface that defines methods for URL storage
type DataStore interface {
	AddURL(urlData URLData) error
	GetURL(shortURL string) (URLData, bool)
	DeleteURL(shortURL string) error
	GetAllURLs() []URLData
}

// URLData represents a shortened URL entry
type URLData struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
