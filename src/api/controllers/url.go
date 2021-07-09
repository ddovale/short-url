package controllers

import (
	"net/http"
	"strconv"

	"github.com/ddovale/short-url/src/api/services/urls"
	"github.com/gin-gonic/gin"
)

func GetAllUrls(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, urls.All())
}

func GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	url := urls.Get(id)
	if url == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Url not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"url": *url,
	})
}

func CreateUrl(ctx *gin.Context) {
	var url urls.Url
	if err := ctx.BindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if url.Link == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Link cannot be empty"})
		return
	}

	urls.Create(&url)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully Created",
		"url":     url,
	})
}
