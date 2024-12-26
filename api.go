package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	secretKey string
	rateLimit = 10
	usageMap  = make(map[string]int)
	mutex     = sync.Mutex{}
	validate  = validator.New()
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/login", loginHandler)

	orders := router.Group("/orders")
	orders.Use(authMiddleware())
	{
		orders.GET("", getOrders)                               // TEST CASE #1
		orders.GET("/:id", getOrderByID)                        // TEST CASE #2
		orders.POST("", createOrder)                            // TEST CASE #3
		orders.PUT("/:id", updateOrder)                         // TEST CASE #4
		orders.POST("/:id/confirm", confirmOrder)               // TEST CASE #5
		orders.POST("/:id/items", addItemToOrder)               // TEST CASE #6
		orders.PUT("/:id/items/:item_id", updateItemInOrder)    // TEST CASE #7
		orders.DELETE("/:id/items/:item_id", deleteItemInOrder) // TEST CASE #8
	}

	router.Run(":8080")
}

func getOrders(c *gin.Context) { // TEST CASE #1
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	if pageStr == "" && limitStr == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Displaying all data",
			"data":    getAllOrders(),
		})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil && pageStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil && limitStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 10 {
		limit = 10
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Displaying paginated data",
		"page":    page,
		"limit":   limit,
		"data":    getPaginatedOrders(page, limit),
	})

}

func getAllOrders() []order { // TEST CASE #1
	return Orders
}

func getPaginatedOrders(page, limit int) []order { // TEST CASE #1
	allOrders := getAllOrders()
	start := (page - 1) * limit
	end := start + limit
	if start > len(allOrders) {
		return []order{}
	}
	if end > len(allOrders) {
		end = len(allOrders)
	}

	return allOrders[start:end]
}

func getOrderByID(c *gin.Context) { // TEST CASE #2
	id := c.Param("id")

	for _, a := range Orders {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{
				"order": a,
				"items": getOrderDetail(a.ID),
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "order not found"})
}

func getOrderDetail(orderID string) []orderItem { // TEST CASE #2
	var OrderItemsByOrderId []orderItem
	for _, a := range OrderItems {
		if a.OrderID == orderID {
			OrderItemsByOrderId = append(OrderItemsByOrderId, a)
		}
	}
	return OrderItemsByOrderId
}

func createOrder(c *gin.Context) {
	var newOrder order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}
	newOrder.ID = fmt.Sprintf("%d", len(Orders)+1)
	newOrder.Status = "Created"
	newOrder.SubTotal = 0
	newOrder.Taxes = 0
	newOrder.Total = 0
	if err := validate.Struct(newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}

	newOrder.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	newOrder.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	Orders = append(Orders, newOrder)
	c.JSON(http.StatusCreated, gin.H{"data": newOrder})
}

func updateOrder(c *gin.Context) {
	orderID := c.Param("id")
	var updatedData order

	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}
	updatedData.ID = orderID
	if err := validate.Struct(updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}
	for i, order := range Orders {
		if order.ID == orderID {
			updatedData.CreatedAt = order.CreatedAt
			updatedData.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			Orders[i] = updatedData
			c.JSON(http.StatusOK, gin.H{"data": Orders[i]})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func confirmOrder(c *gin.Context) {
	orderID := c.Param("id")

	for i, order := range Orders {
		if order.ID == orderID {
			if addresses, err := getUserAddressesById(order.UserID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				Orders[i].BillingAddress = addresses
				Orders[i].ShippingAddress = addresses
			}
			if order.Status == "Pending" {
				Orders[i].Status = "Confirmed"
				Orders[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
				c.JSON(http.StatusOK, gin.H{"message": "Order confirmed", "data": Orders[i]})
			} else if order.Status == "Confirmed" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Order already confirmed"})
			} else if order.Status == "Created" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Order is empty, put some item to this order to confirm"})
			} else if order.Status == "Deleted" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Order have been deleted"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Status"})
			}
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func addItemToOrder(c *gin.Context) {
	orderID := c.Param("id")
	var newItem orderItem

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}
	newItem.OrderID = orderID
	if err := validate.Struct(newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}

	newItem.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	newItem.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	for _, order := range Orders {
		if order.ID == orderID {
			for _, orderItem := range OrderItems {
				if orderItem.OrderID == orderID && orderItem.ItemID == newItem.ItemID {
					c.JSON(http.StatusBadRequest, gin.H{"error": "this item already in your order, try to update instead"})
					return
				}
			}
			OrderItems = append(OrderItems, newItem)
			c.JSON(http.StatusCreated, gin.H{"data": newItem})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func updateItemInOrder(c *gin.Context) {
	orderID := c.Param("id")
	itemID := c.Param("item_id")
	var updatedItem orderItem

	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}
	updatedItem.OrderID = orderID
	updatedItem.ItemID = itemID

	if err := validate.Struct(updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}

	for _, order := range Orders {
		if order.ID == orderID {
			for j, item := range OrderItems {
				if item.OrderID == orderID && item.ItemID == itemID {
					updatedItem.CreatedAt = item.CreatedAt
					updatedItem.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
					OrderItems[j] = updatedItem
					c.JSON(http.StatusOK, gin.H{"data": OrderItems[j]})
					return
				}
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func deleteItemInOrder(c *gin.Context) {
	orderID := c.Param("id")
	itemID := c.Param("item_id")

	for _, order := range Orders {
		if order.ID == orderID {
			for j, item := range OrderItems {
				if item.OrderID == orderID && item.ItemID == itemID {
					OrderItems = append(OrderItems[:j], OrderItems[j+1:]...)
					c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
					return
				}
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func getUserAddressesById(userID string) (string, error) {

	for _, userInfo := range Users {
		if userInfo.ID == userID {
			if len(userInfo.Addresses) > 0 {
				return userInfo.Addresses, nil
			} else {
				return "", fmt.Errorf("add your address before proceed")
			}
		}
	}
	return "", fmt.Errorf("user not found")
}
