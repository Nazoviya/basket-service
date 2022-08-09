package api

import "github.com/Nazoviya/basketService/util"

// declare variables to hold config values taken from app.env.
var (
	config, _               = util.LoadConfig(".")
	minBasketTotal          = config.MinBasketTotal
	minTotalPurchaseInMonth = config.MinTotalPurchaseInMonth
)

// calculates business logic, takes required values and returns float32 array
// which includes discount for different VAT values.
func calculateBusinessLogic(userOrderCount int8, userItemCount int, userBasketTotal, minBasketTotal, userInMonthTotal, minTotalPurchaseInMonth float32) []float32 {
	if userOrderCount+1 == 4 && userBasketTotal > minBasketTotal {
		if userItemCount < 4 {
			if userInMonthTotal < minTotalPurchaseInMonth {
				return []float32{0, .9, .85} // %1 = 0, %8 = 10, %18 = 15
			} else {
				return []float32{.9, .9, .85} // %1 = 10, %8 = 10, %18 = 15
			}
		} else {
			return []float32{.92, .9, .85} // %1 = 8, %8 = 10, %18 = 15
		}
	} else {
		if userInMonthTotal >= minTotalPurchaseInMonth && userItemCount < 4 {
			return []float32{.9, .9, .9} // %1 = 10, %8 = 10, %18 = 10
		} else if userInMonthTotal < minTotalPurchaseInMonth && userItemCount >= 4 {
			return []float32{.92, .92, .92} // %1 = 8, %8 = 8, %18 = 8
		}
	}
	return []float32{0, 0, 0} // %1 = 0, %8 = 0, %18 = 0
}

// calculates discount for each product in the basket after adding
// and deleting products. Returns total price and total VAT of products.
func calculateDiscount(priceList []float32, vatList []int16, discount []float32) (calculatePrice, calculateVAT float32) {
	// loop over every price on the slice to make discount.
	for i := range priceList {
		switch vatList[i] {
		case 1:
			priceList[i] *= discount[0]
		case 8:
			priceList[i] *= discount[1]
		case 18:
			priceList[i] *= discount[2]
		}
	}

	// calculates total price and total VAT, then returns.
	for i, v := range priceList {
		calculatePrice += v
		calculateVAT += priceList[i]*(1+float32(vatList[i])/100) - priceList[i]
	}
	return calculatePrice, calculateVAT
}
