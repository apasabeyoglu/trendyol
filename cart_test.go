package trendyol

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)
	var (
		cart = NewCart()

		categoryFood = Category{ID: 1, CategoryTitle:"Food", Parent: nil}
		categoryElectronics = Category{ID: 2, CategoryTitle:"Electronics", Parent: nil}
		categoryComputer = Category{ID: 3, CategoryTitle:"Computer", Parent: &categoryElectronics}

		productApple = Product{ID:1, Category: categoryFood, Price:10.00, ProductTitle:"Apple"}
		productStrawberry = Product{ ID:2, Category: categoryFood, Price:15.00, ProductTitle:"Strawberry"}
		productLaptop = Product{ID:3, Category: categoryComputer, Price:5000.00, ProductTitle:"Laptop"}
		productDesktop = Product{ID:4, Category: categoryComputer, Price:10000.00, ProductTitle:"Desktop Computer"}
		productHeadphones = Product{ID:5, Category: categoryElectronics, Price:200.00, ProductTitle:"Wireless Headphones"}

		campaign = Campaign{Percentage: 0.1, RequiredQuantity:2, ProductCategory: []string{"Food", "Electronics"}, Code:"Trendyol Gida ve Elektronik Indirimleri"}

		couponSummer = Coupon{Code: "Trendyol Yaz Indirimleri", Percentage: 0.05, MinimumAmount: 400}
		couponExtreme = Coupon{Code: "50.000 TL Uzeri %40 indirim!", Percentage: 0.25, MinimumAmount: 50000}
	)

// Feature : Products can be added to the cart
func TestCart_AddToCart(t *testing.T) {
	/*
		Scenario 1 : Selected product will be added to an empty cart
		Given : Cart is currently empty
		When : User adds 10 apples to his/her cart
		Then : Cart should consist 10 apples
	*/
	cart.AddToCart(productApple,10)

	assert.Equal(t, cart.Orders[productApple.ProductTitle].Quantity, int64(10), "Expected number of apples: 10" )

	/*
		Scenario 2 : Selected product will be added to the current cart
		Given : Cart has the same item that user wants to add
		When : User adds 10 more apples to his/her cart
		Then : Cart should consist 20 apples
	*/
	cart.AddToCart(productApple,10)

	assert.Equal(t, cart.Orders[productApple.ProductTitle].Quantity, int64(20), "Expected number of apples: 20" )

	/*
		Scenario 3 : Selected product will be added to the cart
		Given : Cart doesn't have the same item that user wants to add
		When : User adds 3 desktop computers to his/her cart
		Then : Cart should consist 20 apples and 3 desktop computers
	*/
	cart.AddToCart(productDesktop, 3)

	assert.Equal(t, cart.Orders[productApple.ProductTitle].Quantity, int64(20), "Expected number of apples: 20" )
	assert.Equal(t, cart.Orders[productDesktop.ProductTitle].Quantity, int64(3), "Expected number of desktop computers: 3" )

}

// Feature : Products can be removed from the cart
func TestCart_DeleteFromCart(t *testing.T) {
	/*
		Scenario 1 : Selected product will be removed from users cart
		Given : Cart currently has 20 apples and 3 desktop computers
		When : User wants to remove 5 apples from his/her cart
		Then : Cart should consist 15 apples
	*/
	cart.DeleteFromCart(productApple,5)

	assert.Equal(t, cart.Orders[productApple.ProductTitle].Quantity, int64(15), "Expected number of apples: 15" )

	/*
		Scenario 2 : User tries to remove an unexisting product from his/her cart
		Given : Cart currently has 15 apples and 3 desktop computers
		When : User wants to remove 5 strawberries from his/her cart
		Then : I should return an error message
	*/

	assert.Equal(t, cart.DeleteFromCart(productStrawberry,5), errors.New("No products to delete!"), "No items to delete")

}

// Feature : Total price of the cart should be calculated
func TestCart_CalculateTotal(t *testing.T) {


	/*
		Scenario 1 : User wants to see the total price of his/her cart
		Given : Cart currently has 15 apples and 3 desktop computers
		When : User applies a campaign for food products which gives 10% discount if they add more than 1 food product
		Then : Cart Total should be 30.135 TRY
	*/
	cart.CalculateTotal(campaign)

	assert.Equal(t, cart.CartTotal, 27135.00, "Cart Total should be : 27,135 TRY")

	/*
		Scenario 2 : User wants to see the total price but there is no campaign to use
		Given : Cart currently has 15 apples and 3 desktop computers
		When : There is no campaign applicable
		Then : No campaign discounts will be applied to newly added product, Cart Total should be 27.150TRY
	*/

	var c Campaign
	cart.CalculateTotal(c)

	assert.Equal(t, cart.CartTotal, 30150.00, "Cart Total should be : 30,150TRY")

	/*
		Scenario 3 : User wants to see the total price of his/her cart with having a non-discount product
		Given : Cart currently has 15 apples and 3 desktop computers
		When : User adds 1 strawberry to cart
		Then : No campaign discounts will be applied to newly added product, Cart Total should be 27.150TRY
	*/

	cart.AddToCart(productStrawberry,1)
	cart.CalculateTotal(campaign)

	assert.Equal(t, cart.CartTotal, 27150.00, "Cart Total should be : 27,150TRY")

}

// Feature : Delivery price should be calculated
func TestCart_CalculateDeliveryFee(t *testing.T) {
	/*
		Scenario 1 : User wants to see the total price of delivery
		Given : Cart currently has 15 apples, 3 desktop computers and 1 strawberry
		When : 3 different products mean 3 different deliveries and fee would be increased 1% after each product
		Then : Cart Total should be 30.135 TRY
	*/
	cart.CalculateDeliveryFee()
	assert.Equal(t, cart.DeliveryFee, 46.00, "Delivery Fee should be : 46TRY")
}

// Feature : Total Price with the coupon discount should be calculated
func TestCart_CalculateCouponDiscount(t *testing.T) {
	/*
		Scenario 1 : User will proceed to checkout to see the total price of his/her shopping
		Given : Cart Total is 27.150TRY and Delivery Fee is 46TRY
		When : User applies a valid coupon to have a discount
		Then : Total fee should be 25838.5TRY
	*/
	cart.CalculateCouponDiscount(couponSummer)

	assert.Equal(t, cart.TotalPrice,25838.5,"Total Fee should be : 25,838.5TRY")

	/*
		Scenario 2 : User will proceed to checkout to see the total price of his/her shopping
		Given : Cart Total is 27.150TRY and Delivery Fee is 46TRY
		When : User tries to apply an invalid coupon
		Then : Total fee should be 25838.5TRY
	*/
	cart.CalculateCouponDiscount(couponExtreme)

	assert.Equal(t, cart.TotalPrice,27196.00,"Total Fee should be : 27,196TRY")
}