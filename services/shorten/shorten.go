package shorten

import (
	"log"
	"time"
	"urlshorter/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostCreateShortUrl(c *gin.Context, db *gorm.DB) {
	response := struct {
		Url       string `json:"url"`
		ShortCode string `json:"shortCode"`
	}{}

	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	newUrl := models.ShortUrl{
		Url:         response.Url,
		ShortCode:   response.ShortCode,
		AccessCount: 0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.Create(&newUrl).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		log.Printf("[ERROR] Database creating with error: %s", err.Error())
		return
	}

	c.JSON(201, gin.H{
		"status":  "success",
		"message": "successfully created",
		"data":    newUrl,
	})
}

func GetShortUrl(c *gin.Context, db *gorm.DB) {
	var result models.ShortUrl

	shorturl := c.Param("shorturl")

	if err := db.Where("short_code = ?", shorturl).Find(&result).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		log.Printf("[ERROR] %s", err.Error())
		return
	}

	if err := db.Model(&result).Where("short_code = ?", shorturl).Update("access_count", result.AccessCount+1).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		log.Printf("[ERROR] %s", err.Error())
		return
	}

	c.JSON(200, gin.H{
		"id":        result.ID,
		"url":       result.Url,
		"shortCode": result.ShortCode,
		"createdAt": result.CreatedAt,
		"updatedAt": result.UpdatedAt,
	})
}

func GetShortUrlStats(c *gin.Context, db *gorm.DB) {
	var result models.ShortUrl

	shorturl := c.Param("shorturl")

	if err := db.Where("short_code = ?", shorturl).Find(&result).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		log.Printf("[ERROR] %s", err.Error())
		return
	}

	c.JSON(200, gin.H{
		"id":          result.ID,
		"url":         result.Url,
		"shortCode":   result.ShortCode,
		"createdAt":   result.CreatedAt,
		"updatedAt":   result.UpdatedAt,
		"accessCount": result.AccessCount,
	})
}

func DeleteShortUrl(c *gin.Context, db *gorm.DB) {
	var result models.ShortUrl

	shorturl := c.Param("shorturl")

	if err := db.Where("short_code = ?", shorturl).Find(&result).Delete(&result).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		log.Printf("[ERROR] %s", err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully deleted",
	})
}

func UpdateShortUrl(c *gin.Context, db *gorm.DB) {
	var url models.ShortUrl

	response := struct {
		Url string `json:"url"`
	}{}
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	shorturl := c.Param("shorturl")

	if err := db.Model(&url).Where("short_code = ?", shorturl).Update("url", response.Url).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		log.Printf("[ERROR] %s", err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully updated",
	})
}
