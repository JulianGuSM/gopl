package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/23 3:17 下午
 * @Desc:
 */

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Red Train", Artist: "Gery Coltrane", Price: 17.99},
	{ID: "3", Title: "White Train", Artist: "Sarah Coltrane", Price: 56.99},
}

var logger, _ = zap.NewProduction()

func getAlbums1(c *gin.Context) {
	start := time.Now()
	c.JSON(http.StatusOK, albums)
	logger.Info("Context.JSON() Time consume...",
		zap.Any("cost:", time.Since(start)),
	)
}

func getAlbums2(c *gin.Context) {
	start := time.Now()
	c.IndentedJSON(http.StatusOK, albums)
	logger.Info("Context.IndentedJSON() Time consume...",
		zap.Any("cost:", time.Since(start)),
	)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	r := gin.Default()

	r.GET("/Albums", getAlbums1)
	r.GET("/Album's", getAlbums2)
	r.POST("/albums", postAlbums)
	r.GET("/albums/:id", getAlbumById)
	r.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	err := r.Run(":9080")
	if err != nil {
		logger.Error("server start failed!",
			zap.Any("error_msg", err),
		)
		return
	}
}
