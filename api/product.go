package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create a struct to store properties of products.
type getProductsRequest struct {
	ProductID    int64  `form:"product_id"`
	ProductName  string `form:"product_name"`
	ProductPrice string `form:"product_price"`
	ProductVat   int16  `form:"product_vat"`
}

// gets all products available on database.
func (server *Server) getProducts(ctx *gin.Context) {
	var req getProductsRequest
	// bind query to store all products.
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// store var v as []db.Product object via calling ListProducts
	// function with ctx *gin.Context pointer. Returns error if nil.
	v, err := server.store.ListProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// check if v is nil, posts context or error message.
	if v != nil {
		ctx.JSON(http.StatusOK, v)
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"msg": "no product is available."})
	}
}
