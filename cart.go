package trendyol

import "errors"

type Cart struct {
	ID          int64
	Orders      map[string]Order
	CartTotal   float64
	DeliveryFee float64
	TotalPrice  float64
}

type Order struct {
	OrderProduct Product
	Quantity     int64
}

func NewCart() *Cart {
	m := new(Cart)
	m.Orders = map[string]Order{}
	return m
}

func (c *Cart) AddToCart(product Product, qty int64) {
	order, ok := c.Orders[product.ProductTitle]
	if !ok {
		order.OrderProduct = product
		order.Quantity = qty
	} else {
		order.Quantity += qty
	}
	c.Orders[product.ProductTitle] = order
}

func (c *Cart) DeleteFromCart(product Product, qty int64) error {
	order, ok := c.Orders[product.ProductTitle]
	if ok {
		order.Quantity -= qty
		c.Orders[product.ProductTitle] = order
		return nil
	}
	return errors.New("No products to delete!")
}

func (c *Cart) CalculateTotal(campaign Campaign) {
	price := 0.0
	for _, order := range c.Orders {
		if campaign.ValidForCategory(order.OrderProduct.Category) && order.Quantity >= campaign.RequiredQuantity {
			price += (order.OrderProduct.Price * (1 - campaign.Percentage)) * float64(order.Quantity)
		} else {
			price += order.OrderProduct.Price * float64(order.Quantity)
		}

	}
	c.CartTotal = price

}

// I'm assuming that every order would be a different delivery. I will use 10 TRY as a fixed amount for 1 product and increase it by 10% for each product after the first one
func (c *Cart) CalculateDeliveryFee() {
	deliveryFee := 0.0

	for _, order := range c.Orders {
		deliveryFee += 10.0 + (10.0 * 0.1 * float64(order.Quantity-1))
	}

	c.DeliveryFee = deliveryFee
}

func (c *Cart) CalculateCouponDiscount(coupon Coupon) {
	if c.CartTotal >= coupon.MinimumAmount {
		c.TotalPrice = c.CartTotal + c.DeliveryFee - (c.CartTotal * coupon.Percentage)
	} else {
		c.TotalPrice = c.CartTotal + c.DeliveryFee
	}
}
