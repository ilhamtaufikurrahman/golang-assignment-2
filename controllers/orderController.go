package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	var newOrder models.Order

	err := ctx.ShouldBindJSON(&newOrder)
	fmt.Println(newOrder)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Create(&newOrder).Error

	if err != nil {
		fmt.Println("Error creating data:", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

func GetOrders(ctx *gin.Context) {
	db := database.GetDB()

	var orders []models.Order

	dberr := db.Preload("Items").Find(&orders).Error

	if dberr != nil {
		panic(dberr)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}
