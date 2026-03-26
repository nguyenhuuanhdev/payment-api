package services

import "payment-api/models"

var orders []models.Order
var currentID = 1

//Tạo đơn hàng mới
func CreateOrder(amount float64) models.Order {
	order := models.Order{
		ID:     currentID,
		Amount: amount,
		Status: "pending",
	}
	currentID++
	orders = append(orders, order)
	return order
}

//Lấy đơn hàng theo Id
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
