package models

import "errors"

// OrderDTO represents a user's order information
type OrderDTO struct {
	UserID      string      `json:"userId"`
	ProductList []Product   `json:"productList"`
	PaymentInfo PaymentInfo `json:"paymentInfo"`
}

// Product represents a product in an order
type Product struct {
	ProductID   string  `json:"productId"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
}

// PaymentInfo represents the payment details for an order
type PaymentInfo struct {
	PaymentMethod string  `json:"paymentMethod"`
	Amount        float64 `json:"amount"`
}

// ValidateOrderDTO validates the OrderDTO struct
func ValidateOrderDTO(orderDTO OrderDTO) error {
	if orderDTO.UserID == "" {
		return errors.New("user ID cannot be empty")
	}

	if len(orderDTO.ProductList) == 0 {
		return errors.New("order must have at least one product")
	}

	for _, product := range orderDTO.ProductList {
		if product.ProductID == "" {
			return errors.New("product ID cannot be empty")
		}

		if product.ProductName == "" {
			return errors.New("product name cannot be empty")
		}

		if product.Quantity <= 0 {
			return errors.New("product quantity must be greater than zero")
		}

		if product.UnitPrice <= 0 {
			return errors.New("product unit price must be greater than zero")
		}
	}

	if orderDTO.PaymentInfo.PaymentMethod == "" {
		return errors.New("payment method cannot be empty")
	}

	if orderDTO.PaymentInfo.Amount <= 0 {
		return errors.New("payment amount must be greater than zero")
	}

	return nil
}
