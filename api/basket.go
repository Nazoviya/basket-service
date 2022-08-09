package api

import (
	"database/sql"
	"net/http"

	db "github.com/Nazoviya/basketService/db/sqlc"
	"github.com/gin-gonic/gin"
)

// create a struct to store properties of products.
type addToBasketRequest struct {
	ProductID    int64  `uri:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice string `json:"product_price"`
	ProductVat   int16  `json:"product_vat"`
}

// adds specified product to the basket.
func (server *Server) addToBasket(ctx *gin.Context) {
	var req addToBasketRequest
	// bind URI to specify product to add.
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// store var v as db.Userbasket object via calling AddToBasket
	// function with ctx *gin.Context pointer and specified product_id.
	v, err := server.store.AddToBasket(ctx, req.ProductID)
	if err != nil {
		// Returns no sql row error.
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		// returns internal server error, in gin error type.
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// serializes the given struct v as JSON into the response body.
	ctx.JSON(http.StatusOK, v)
}

// create a struct to store properties of products.
type showBasketRequest struct {
	ProductID    int64  `form:"product_id"`
	ProductName  string `form:"product_name"`
	ProductPrice string `form:"product_price"`
	ProductVat   int16  `form:"product_vat"`
}

// shows basket.
func (server *Server) showBasket(ctx *gin.Context) {
	var req showBasketRequest
	// bind query to show all products.
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// store var v as []db.Userbasket object via calling ShowBasket
	// function with ctx *gin.Context pointer. Returns error if nil.
	v, err := server.store.ShowBasket(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// check if v is nil, posts context or error message.
	if v != nil {
		ctx.JSON(http.StatusOK, v)
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"msg": "basket is empty."})
	}
}

// create a struct to store properties of products.
type deleteFromBasketRequest struct {
	ProductID    int64  `uri:"product_id"`
	ProductName  string `form:"product_name"`
	ProductPrice string `form:"product_price"`
	ProductVat   int16  `form:"product_vat"`
}

// deletes specified product from basket.
func (server *Server) deleteFromBasket(ctx *gin.Context) {
	var req deleteFromBasketRequest
	// bind URI to specify product to delete.
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// quantity issue for feature usage.
	// arg := db.DeleteFromBasketParams{
	// 	UserID:    req.UserID,
	// 	ProductID: req.ProductID,
	// }

	// store var v as error via calling DeleteFromBasket function
	// with ctx *gin.Context pointer specified product_id.
	v := server.store.DeleteFromBasket(ctx, req.ProductID)

	// check if v is nil, posts context or error message.
	if v != nil {
		ctx.JSON(http.StatusOK, v)
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"msg": "item is deleted from basket."})
	}
}

// create a struct to store properties of basket_total.
type calculateBasketRequest struct {
	Price      float32 `json:"price"`
	Vat        float32 `json:"vat"`
	TotalPrice float32 `json:"total_price"`
	Discount   float32 `json:"discount"`
}

// calculates basket_total and posts as JSON object.
func (server *Server) calculateBasket(ctx *gin.Context) {
	var req calculateBasketRequest
	// bind query to show all products.
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// holds discountDriver functions values.
	reqPrice, reqVAT, reqTotalPrice, reqBasketDiscount := discountDriver(server, ctx)

	// arguments for posting required values.
	arg := db.CalculateBasketParams{
		Price:      reqPrice,
		Vat:        reqVAT,
		TotalPrice: reqTotalPrice - reqBasketDiscount,
		Discount:   reqBasketDiscount,
	}

	// store var v as db.TotalBasket via calling CalculateBasket function
	// with ctx *gin.Context pointer and db.CalculateBasketParams created.
	v, err := server.store.CalculateBasket(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// serializes the given struct v as JSON into the response body.
	ctx.JSON(http.StatusOK, v)
}

// driver function to call created functions with provided default values.
func discountDriver(server *Server, ctx *gin.Context) (reqPrice, reqVAT, reqTotalPrice, reqBasketDiscount float32) {
	reqPrice = calculatePrice(server, ctx)
	reqVAT = calculateVAT(server, ctx)
	reqTotalPrice = reqPrice + reqVAT

	// default user values
	var userBasketTotal float32 = 100
	var userOrderCount int8 = 3
	var userInMonthTotal float32 = 1000
	var userItemCount int = 1

	// holds prices as []float32 slice for all products in the basket.
	priceList, err := server.store.GetPrice(ctx)
	if err != nil {
		errorResponse(err)
		return
	}

	// holds VATs as []float32 slice for all products in the basket.
	vatList, err := server.store.GetVAT(ctx)
	if err != nil {
		errorResponse(err)
		return
	}

	// holds discount values for every product.
	var discount []float32 = calculateBusinessLogic(userOrderCount, userItemCount, userBasketTotal, minBasketTotal, userInMonthTotal, minTotalPurchaseInMonth)

	// calculates discount with values given, returns total price and VAT.
	disreqPrice, disreqVAT := calculateDiscount(priceList, vatList, discount)

	disReqTotalPrice := disreqPrice + disreqVAT
	reqBasketDiscount = reqTotalPrice - disReqTotalPrice

	if disReqTotalPrice == 0 {
		reqBasketDiscount = disReqTotalPrice
	}
	return reqPrice, reqVAT, reqTotalPrice, reqBasketDiscount
}

// calculates total price of the basket.
func calculatePrice(server *Server, ctx *gin.Context) float32 {
	priceList, err := server.store.GetPrice(ctx)
	if err != nil {
		errorResponse(err)
	}

	var calculatePrice float32
	for _, v := range priceList {
		calculatePrice += v
	}
	return calculatePrice
}

// calculates total VAT of the basket.
func calculateVAT(server *Server, ctx *gin.Context) float32 {
	priceList, err := server.store.GetPrice(ctx)
	if err != nil {
		errorResponse(err)
	}
	vatList, err := server.store.GetVAT(ctx)
	if err != nil {
		errorResponse(err)
	}

	var calculateVAT float32
	for i := range priceList {
		calculateVAT += priceList[i]*(1+float32(vatList[i])/100) - priceList[i]
	}
	return calculateVAT
}
