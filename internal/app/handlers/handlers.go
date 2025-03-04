package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/services"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg     *config.Config
	Service services.URLShortener
}

func NewHandler(cfg *config.Config, service services.URLShortener) *Handler {
	return &Handler{cfg: cfg, Service: service}
}

func (h *Handler) URLCreator(c *gin.Context) {
	defer c.Request.Body.Close()

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	text, err := url.Parse(string(body))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Empty body"})
		return
	}

	shortURL, err := h.Service.ShortenURL(text.String())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusCreated, h.cfg.BaseURL+"/"+shortURL)
}

func (h *Handler) URLCreatorJSON(c *gin.Context) {
	defer c.Request.Body.Close()

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var req struct {
		URL string `json:"url"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	shortURL, err := h.Service.ShortenURL(req.URL)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := struct {
		Result string `json:"result"`
	}{
		Result: h.cfg.BaseURL + "/" + shortURL,
	}

	c.Header("Content-Type", "application/json")
	c.Status(http.StatusCreated)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(resp); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode response"})
	}
}

func (h *Handler) GetURL(c *gin.Context) {
	key := c.Param("url")
	originalURL, exists := h.Service.GetOriginalURL(key)
	if !exists {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Header("Location", originalURL)
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusTemporaryRedirect, originalURL)
}
