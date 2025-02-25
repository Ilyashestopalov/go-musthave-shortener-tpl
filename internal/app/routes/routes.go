package routes

import (
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/interfaces"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/shortner"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(router *gin.Engine, shortener interfaces.URLShortener, baseURL string) {
	router.POST("/api/shorten", shortner.APIShortenURLHandler(shortener, baseURL))
	router.POST("/", shortner.ShortenURLHandler(shortener, baseURL))
	router.GET("/:shortened", shortner.RedirectURLHandler(shortener))
}
