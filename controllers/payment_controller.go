package controllers

import (
	"net/http"
	"payment-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var input struct {
		Amount float64 `json: "amount"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dữ liệu không hợp lệ",
		})
		return
	}

	order := services.CreateOrder(input.Amount)
	c.JSON(http.StatusOK, gin.H{
		"message": "Tạo thanh toán thành công",
		"order":   order,
	})
}

// Get id
func GetOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID không hợp lệ",
		})
		return
	}
	order, found := services.GetOrderByID(id)
	if !found {
		c.JSON(404, gin.H{"error": "Không tìm thấy đơn hàng"})
		return
	}
	c.JSON(200, gin.H{
		"order": order,
	})
}

// Get /pay/:id
func Pay(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID không hợp lệ",
		})
		return
	}

	order, found := services.PayOrder(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Đơn hàng không tồn tại",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Thanh toán thành công",
		"order":   order,
	})
}

func UpdatePayment(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Dữ liệu không hợp lệ",
		})
		return
	}
	order, found := services.UpdateOrder(id, input.Amount)

	if !found {
		c.JSON(404, gin.H{
			"error": "Không tìm thấy đơn",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Cập nhật thành công",
		"order":   order,
	})
}

func DeleteOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	deleted := services.DeleteOrder(id)
	if !deleted {
		c.JSON(404, gin.H{"error": "Không tìm thấy đơn"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Xóa thành công",
	})
}
