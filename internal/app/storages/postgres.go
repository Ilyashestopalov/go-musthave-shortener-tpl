package storages

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// PostgreSQLStore implements DataStore and holds URLs in a PostgreSQL database
type PostgreSQLStore struct {
	DB *sqlx.DB
}

// NewPostgreSQLStore creates a new PostgreSQLStore
func NewPostgreSQLStore(dataSourceName string) (*PostgreSQLStore, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &PostgreSQLStore{DB: db}, nil
}

// AddURL adds a URL to the store
func (pg *PostgreSQLStore) AddURL(urlData URLData) error {
	_, err := pg.DB.Exec("INSERT INTO urls (uuid, short_url, original_url) VALUES ($1, $2, $3)",
		urlData.UUID, urlData.ShortURL, urlData.OriginalURL)
	return err
}

// GetURL retrieves a URL by short URL
func (pg *PostgreSQLStore) GetURL(shortURL string) (URLData, bool) {
	var urlData URLData
	err := pg.DB.Get(&urlData, "SELECT uuid, short_url, original_url FROM urls WHERE short_url = $1", shortURL)
	if err != nil {
		return URLData{}, false
	}
	return urlData, true
}

// DeleteURL deletes a URL from the store
func (pg *PostgreSQLStore) DeleteURL(shortURL string) error {
	_, err := pg.DB.Exec("DELETE FROM urls WHERE short_url = $1", shortURL)
	return err
}

// GetAllURLs returns all URLs
func (pg *PostgreSQLStore) GetAllURLs() []URLData {
	var urls []URLData
	err := pg.DB.Select(&urls, "SELECT uuid, short_url, original_url FROM urls")
	if err != nil {
		return nil
	}
	return urls
}
