package controllers

import (
	"assignment-2/database"
	"assignment-2/helpers"
	"assignment-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func CreateOrder(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Order := models.Order{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Order)
	} else {
		c.ShouldBind(&Order)
	}

	err := db.Debug().Create(&Order).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Order)

}

func GetOrders(c *gin.Context) {
	db := database.GetDB()
	Orders := []models.Order{}

	err := db.Preload("Items").Find(&Orders).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Orders)
}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()
	orderID, err := strconv.Atoi(c.Param("OrderID"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Invalid param orderId",
		})
		return
	}

	err = db.Delete(models.Order{}, "order_id", orderID).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error deleting",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete order",
	})
}

func UpdateOrder(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	OrderID, err := strconv.Atoi(c.Param("OrderID"))
	UpdateOrder := models.Order{}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Invalid param orderId",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&UpdateOrder)
	} else {
		c.ShouldBind(&UpdateOrder)
	}

	for i := range UpdateOrder.Items {
		err = db.Model(&UpdateOrder.Items[i]).Where("item_id=?", UpdateOrder.Items[i].ItemID).Updates(&UpdateOrder.Items[i]).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Error updating item",
				"message": err.Error(),
			})
			return
		}
	}

	err = db.Model(&UpdateOrder).Where("order_id=?", OrderID).Omit("Items").Updates(&UpdateOrder).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error updating order",
			"message": err.Error(),
		})
		return
	}

	err = db.Preload("Items").Where("order_id=?", OrderID).Find(&UpdateOrder).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfuly update data",
		"data":    UpdateOrder,
	})
}

// func UpdateItem(c *gin.Context) {
// 	db := database.GetDB()
// 	contentType := helpers.GetContentType(c)
// 	itemID, err := strconv.Atoi(c.Param("ItemID"))
// 	UpdateItem := models.Item{}

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"code":    "500",
// 			"message": "Invalid param orderId",
// 		})
// 		return
// 	}

// 	if contentType == appJSON {
// 		c.ShouldBindJSON(&UpdateItem)
// 	} else {
// 		c.ShouldBind(&UpdateItem)
// 	}

// 	err = db.Model(&UpdateItem).Where("item_id=?", itemID).Updates(&UpdateItem).Error

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"error":   "Error updating",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, UpdateItem)
// }
