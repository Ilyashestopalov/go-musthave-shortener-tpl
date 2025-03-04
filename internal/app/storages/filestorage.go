package storages

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"sync"
)

type URLRecord struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type FileStorage struct {
	mu       sync.RWMutex
	filePath string
	data     map[string]string
}

func NewFileStore(filePath string) *FileStorage {
	store := &FileStorage{
		filePath: filePath,
		data:     make(map[string]string),
	}
	store.loadFromFile()
	return store
}

func (fs *FileStorage) loadFromFile() {
	//fmt.Printf(fs.filePath)
	file, err := os.Open(fs.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var record URLRecord
		if err := json.Unmarshal(scanner.Bytes(), &record); err == nil {
			fs.data[record.ShortURL] = record.OriginalURL
		}
	}
}

func (fs *FileStorage) saveToFile(shortURL, originalURL string) {
	//fmt.Printf(fs.filePath)
	file, err := os.OpenFile(fs.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	record := URLRecord{
		UUID:        shortURL,
		ShortURL:    shortURL,
		OriginalURL: originalURL,
	}

	jsonData, _ := json.Marshal(record)
	file.Write(jsonData)
	file.Write([]byte("\n"))
}

func (fs *FileStorage) Get(shortURL string) (string, bool) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	value, exists := fs.data[shortURL]
	return value, exists
}

func (fs *FileStorage) Set(shortURL, originalURL string) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.data[shortURL] = originalURL
	fs.saveToFile(shortURL, originalURL)
}
