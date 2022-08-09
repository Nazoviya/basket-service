package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create a struct to store properties of products.
type completeOrderRequest struct {
	ProductID    int64  `form:"product_id"`
	ProductName  string `form:"product_name"`
	ProductPrice string `form:"product_price"`
	ProductVat   int16  `form:"product_vat"`
}

// returns completion of order and clears basket.
func (server *Server) completeOrder(ctx *gin.Context) {
	var req completeOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// store var v as error calling CompleteOrder function with
	// ctx *gin.Context pointer.
	v := server.store.CompleteOrder(ctx)

	// check if v is nil, posts context or error message.
	if v != nil {
		ctx.JSON(http.StatusOK, v)
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"msg": "order completed."})
	}
}
