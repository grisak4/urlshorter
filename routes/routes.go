package routes

import (
	"urlshorter/services/shorten"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/shorten", func(ctx *gin.Context) {
		shorten.PostCreateShortUrl(ctx, db)
	})

	r.GET("/shorten/:shorturl", func(ctx *gin.Context) {
		shorten.GetShortUrl(ctx, db)
	})
	r.GET("/shorten/:shorturl/stats", func(ctx *gin.Context) {
		shorten.GetShortUrlStats(ctx, db)
	})

	r.DELETE("/shorten/:shorturl", func(ctx *gin.Context) {
		shorten.DeleteShortUrl(ctx, db)
	})

	r.PUT("/shorten/:shorturl", func(ctx *gin.Context) {
		shorten.UpdateShortUrl(ctx, db)
	})
}
