package main

type user struct {
	ID        string `json:"id,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password,omitempty"`
	Addresses string `json:"addresses,omitempty"`
	CreatedAt string `json:"create_at,omitempty"`
	UpdatedAt string `json:"update_at,omitempty"`
}

var Users = []user{
	{ID: "1", Role: "Admin", Email: "admin@abc.com", Username: "admin", Password: "$2a$10$us2z9KYv/FbpDhob4oETvOO3LLn2NyO6u3/Hr5b430v/JHtKHPDAa", Addresses: "751 Oak St, Townsville", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{ID: "2", Role: "Vendor", Email: "vendor@abc.com", Username: "vendor", Password: "$2a$10$us2z9KYv/FbpDhob4oETvOO3LLn2NyO6u3/Hr5b430v/JHtKHPDAa", Addresses: "751 Oak St, Townsville", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{ID: "3", Role: "Costumer", Email: "a@abc.com", Username: "aaa", Password: "$2a$10$us2z9KYv/FbpDhob4oETvOO3LLn2NyO6u3/Hr5b430v/JHtKHPDAa", Addresses: "751 Oak St, Townsville", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{ID: "4", Role: "Costumer", Email: "b@abc.com", Username: "bbb", Password: "$2a$10$us2z9KYv/FbpDhob4oETvOO3LLn2NyO6u3/Hr5b430v/JHtKHPDAa", Addresses: "751 Oak St, Townsville", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{ID: "5", Role: "Costumer", Email: "c@abc.com", Username: "ccc", Password: "$2a$10$us2z9KYv/FbpDhob4oETvOO3LLn2NyO6u3/Hr5b430v/JHtKHPDAa", Addresses: "751 Oak St, Townsville", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
}

type order struct {
	ID              string  `json:"id" validate:"required"`
	UserID          string  `json:"user_id" validate:"required"`
	Status          string  `json:"status" validate:"required"`
	SubTotal        float64 `json:"sub_total"`
	Taxes           float64 `json:"taxes"`
	Total           float64 `json:"total"`
	BillingAddress  string  `json:"billing_address,omitempty"`
	ShippingAddress string  `json:"shipping_address,omitempty"`
	CreatedAt       string  `json:"create_at,omitempty"`
	UpdatedAt       string  `json:"update_at,omitempty"`
}

var Orders = []order{
	{ID: "1", UserID: "3", Status: "Created", SubTotal: 249.67, Taxes: 46.84, Total: 296.51, BillingAddress: "751 Oak St, Townsville", ShippingAddress: "600 Main St, Townsville", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{ID: "2", UserID: "4", Status: "Created", SubTotal: 286.14, Taxes: 36.21, Total: 322.34, BillingAddress: "904 Main St, Hamletland", ShippingAddress: "156 Elm St, Hamletland", CreatedAt: "2024-10-21 19:52:40", UpdatedAt: "2024-09-17 19:52:40"},
	{ID: "3", UserID: "3", Status: "Created", SubTotal: 103.66, Taxes: 29.74, Total: 133.4, BillingAddress: "925 Oak St, Metropolis", ShippingAddress: "957 Elm St, Hamletland", CreatedAt: "2024-12-10 19:52:40", UpdatedAt: "2024-12-25 19:52:40"},
	{ID: "4", UserID: "4", Status: "Created", SubTotal: 318.36, Taxes: 37.37, Total: 355.73, BillingAddress: "260 Maple Ave, Metropolis", ShippingAddress: "130 Oak St, Townsville", CreatedAt: "2024-12-09 19:52:40", UpdatedAt: "2024-12-10 19:52:40"},
	{ID: "5", UserID: "3", Status: "Created", SubTotal: 467.68, Taxes: 14.18, Total: 481.86, BillingAddress: "481 Main St, Cityville", ShippingAddress: "594 Main St, Townsville", CreatedAt: "2024-12-11 19:52:40", UpdatedAt: "2024-12-03 19:52:40"},
	{ID: "6", UserID: "4", Status: "Pending", SubTotal: 106.83, Taxes: 18.26, Total: 125.09, BillingAddress: "528 Oak St, Villagetown", ShippingAddress: "491 Main St, Hamletland", CreatedAt: "2024-11-17 19:52:40", UpdatedAt: "2024-09-29 19:52:40"},
	{ID: "7", UserID: "3", Status: "Pending", SubTotal: 388.09, Taxes: 17.03, Total: 405.12, BillingAddress: "920 Oak St, Villagetown", ShippingAddress: "315 Maple Ave, Hamletland", CreatedAt: "2024-09-19 19:52:40", UpdatedAt: "2024-09-17 19:52:40"},
	{ID: "8", UserID: "4", Status: "Pending", SubTotal: 405.4, Taxes: 24.79, Total: 430.19, BillingAddress: "726 Maple Ave, Villagetown", ShippingAddress: "418 Pine St, Hamletland", CreatedAt: "2024-09-19 19:52:40", UpdatedAt: "2024-10-02 19:52:40"},
	{ID: "9", UserID: "3", Status: "Pending", SubTotal: 314.02, Taxes: 19.07, Total: 333.09, BillingAddress: "638 Maple Ave, Villagetown", ShippingAddress: "966 Pine St, Metropolis", CreatedAt: "2024-11-28 19:52:40", UpdatedAt: "2024-11-16 19:52:40"},
	{ID: "10", UserID: "4", Status: "Pending", SubTotal: 197.88, Taxes: 14.1, Total: 211.98, BillingAddress: "545 Main St, Villagetown", ShippingAddress: "984 Pine St, Townsville", CreatedAt: "2024-09-19 19:52:40", UpdatedAt: "2024-09-16 19:52:40"},
	{ID: "11", UserID: "3", Status: "Confirmed", SubTotal: 473.71, Taxes: 15.65, Total: 489.35, BillingAddress: "954 Elm St, Hamletland", ShippingAddress: "484 Maple Ave, Metropolis", CreatedAt: "2024-09-28 19:52:40", UpdatedAt: "2024-10-01 19:52:40"},
	{ID: "12", UserID: "4", Status: "Confirmed", SubTotal: 348.55, Taxes: 19.16, Total: 367.71, BillingAddress: "868 Main St, Hamletland", ShippingAddress: "715 Maple Ave, Townsville", CreatedAt: "2024-09-26 19:52:40", UpdatedAt: "2024-09-26 19:52:40"},
	{ID: "13", UserID: "3", Status: "Confirmed", SubTotal: 88.85, Taxes: 46.81, Total: 135.66, BillingAddress: "514 Oak St, Hamletland", ShippingAddress: "506 Oak St, Townsville", CreatedAt: "2024-10-18 19:52:40", UpdatedAt: "2024-11-23 19:52:40"},
	{ID: "14", UserID: "4", Status: "Confirmed", SubTotal: 213.51, Taxes: 25.79, Total: 239.29, BillingAddress: "148 Elm St, Metropolis", ShippingAddress: "373 Maple Ave, Cityville", CreatedAt: "2024-11-06 19:52:40", UpdatedAt: "2024-09-19 19:52:40"},
	{ID: "15", UserID: "3", Status: "Confirmed", SubTotal: 300.06, Taxes: 8.28, Total: 308.34, BillingAddress: "276 Oak St, Townsville", ShippingAddress: "957 Maple Ave, Metropolis", CreatedAt: "2024-11-11 19:52:40", UpdatedAt: "2024-12-20 19:52:40"},
	{ID: "16", UserID: "4", Status: "Deleted", SubTotal: 485.06, Taxes: 43.81, Total: 528.87, BillingAddress: "459 Elm St, Cityville", ShippingAddress: "806 Elm St, Townsville", CreatedAt: "2024-11-12 19:52:40", UpdatedAt: "2024-09-20 19:52:40"},
	{ID: "17", UserID: "3", Status: "Deleted", SubTotal: 129.48, Taxes: 44.74, Total: 174.22, BillingAddress: "472 Main St, Villagetown", ShippingAddress: "750 Main St, Cityville", CreatedAt: "2024-11-17 19:52:40", UpdatedAt: "2024-10-24 19:52:40"},
	{ID: "18", UserID: "4", Status: "Deleted", SubTotal: 422.53, Taxes: 26.44, Total: 448.96, BillingAddress: "944 Pine St, Villagetown", ShippingAddress: "904 Oak St, Villagetown", CreatedAt: "2024-10-13 19:52:40", UpdatedAt: "2024-11-19 19:52:40"},
	{ID: "19", UserID: "3", Status: "Deleted", SubTotal: 239.72, Taxes: 44.91, Total: 284.63, BillingAddress: "361 Elm St, Villagetown", ShippingAddress: "336 Maple Ave, Hamletland", CreatedAt: "2024-11-14 19:52:40", UpdatedAt: "2024-12-07 19:52:40"},
	{ID: "20", UserID: "4", Status: "Deleted", SubTotal: 294.22, Taxes: 21.28, Total: 315.5, BillingAddress: "352 Elm St, Townsville", ShippingAddress: "595 Pine St, Villagetown", CreatedAt: "2024-12-12 19:52:40", UpdatedAt: "2024-10-28 19:52:40"},
}

type inventory struct {
	ID          string  `json:"id" validate:"required"`
	VendorID    string  `json:"vendor_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description,omitempty"`
	Stock       int     `json:"stock" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	CreatedAt   string  `json:"create_at,omitempty"`
	UpdatedAt   string  `json:"update_at,omitempty"`
}

var Inventory = []inventory{
	{ID: "1", VendorID: "2", Name: "A", Description: "JustItem", Stock: 46, Price: 296.51, CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{ID: "2", VendorID: "2", Name: "B", Description: "JustItem", Stock: 36, Price: 322.34, CreatedAt: "2024-10-21 19:52:40", UpdatedAt: "2024-09-17 19:52:40"},
	{ID: "3", VendorID: "2", Name: "C", Description: "JustItem", Stock: 29, Price: 133.40, CreatedAt: "2024-12-10 19:52:40", UpdatedAt: "2024-12-25 19:52:40"},
	{ID: "4", VendorID: "2", Name: "D", Description: "JustItem", Stock: 37, Price: 355.73, CreatedAt: "2024-12-09 19:52:40", UpdatedAt: "2024-12-10 19:52:40"},
	{ID: "5", VendorID: "2", Name: "E", Description: "JustItem", Stock: 14, Price: 481.86, CreatedAt: "2024-12-11 19:52:40", UpdatedAt: "2024-12-03 19:52:40"},
	{ID: "6", VendorID: "2", Name: "F", Description: "JustItem", Stock: 18, Price: 125.09, CreatedAt: "2024-11-17 19:52:40", UpdatedAt: "2024-09-29 19:52:40"},
	{ID: "7", VendorID: "2", Name: "G", Description: "JustItem", Stock: 17, Price: 405.12, CreatedAt: "2024-09-19 19:52:40", UpdatedAt: "2024-09-17 19:52:40"},
	{ID: "8", VendorID: "2", Name: "H", Description: "JustItem", Stock: 24, Price: 430.19, CreatedAt: "2024-09-19 19:52:40", UpdatedAt: "2024-10-02 19:52:40"},
	{ID: "9", VendorID: "2", Name: "I", Description: "JustItem", Stock: 19, Price: 333.09, CreatedAt: "2024-11-28 19:52:40", UpdatedAt: "2024-11-16 19:52:40"},
	{ID: "10", VendorID: "2", Name: "J", Description: "JustItem", Stock: 14, Price: 211.98, CreatedAt: "2024-09-19 19:52:40", UpdatedAt: "2024-09-16 19:52:40"},
}

type orderItem struct {
	OrderID   string `json:"order_id" validate:"required"`
	ItemID    string `json:"item_id" validate:"required"`
	Quantity  string `json:"quantity" validate:"required"`
	Price     string `json:"price" validate:"required"`
	CreatedAt string `json:"create_at,omitempty"`
	UpdatedAt string `json:"update_at,omitempty"`
}

var OrderItems = []orderItem{
	{OrderID: "1", ItemID: "1", Quantity: "1", Price: "123", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{OrderID: "1", ItemID: "2", Quantity: "2", Price: "234", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{OrderID: "2", ItemID: "1", Quantity: "1", Price: "132", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{OrderID: "2", ItemID: "3", Quantity: "2", Price: "232", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
	{OrderID: "3", ItemID: "4", Quantity: "1", Price: "242", CreatedAt: "2024-10-08 19:52:40", UpdatedAt: "2024-09-23 19:52:40"},
}
