package services

import "payment-api/models"

var orders []models.Order
var currentID = 1

//Tạo đơn hàng mới
//Tạo đơn hàng mới với số tiền và trạng thái mặc định là "pending"

//Lấy đơn hàng theo Id 1  2 3
func GetOrderByID(id int) (*models.Order, bool) {
	for _, order := range orders {
		if order.ID == id {
			return &order, true
		}
	}
	return nil, false
}

//Lấy tất cả đơn hàng

func GetAllOrders() []models.Order {
	return orders
}
func GetOrdersByStatus(status string) []models.Order {
	var filteredOrders []models.Order
	for _, order := range orders {
		if order.Status == status {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}
//Thanh toán đơn hàng
func PayOrder(id int) (*models.Order, bool) {
	for i, order := range orders {
		if order.ID == id {
			orders[i].Status = "paid"
			return &orders[i], true
		}
	}
	return nil, false
}

//Sửa đơn hàng
func UpdateOrder(id int, amount float64) (*models.Order, bool) {
	for i, order := range orders {
		if order.ID == id {
			orders[i].Amount = amount
			return &orders[i], true
		}
	}
	return nil, false
}

//Xóa đơn hàng
func DeleteOrder(id int) bool {
	for i, order := range orders {
		if order.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			return true
		}
	}
	return false
}
