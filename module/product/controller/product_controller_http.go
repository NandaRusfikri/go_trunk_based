package controller

import (
	"github.com/gin-gonic/gin"
	"go_trunk_based/dto"
	"go_trunk_based/module/product"
	"go_trunk_based/util"
	"net/http"
)

type ProductControllerHTTP struct {
	productUsecase product.UseCaseInterface
}

func InitProductControllerHTTP(route *gin.Engine, service product.UseCaseInterface) {

	controller := &ProductControllerHTTP{
		productUsecase: service,
	}
	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/products", controller.ProductList)
}

// ProductList
//
//	@Tags			Product
//	@Summary		Product List
//	@Description	Product List
//	@ID				Item-GetList
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			request query	dto.ProductsRequest true "as"
//	@Success		200
//	@Router			/v1/products [get]
func (c *ProductControllerHTTP) ProductList(ctx *gin.Context) {

	var input dto.ProductsRequest

	if err := ctx.ShouldBindQuery(&input); err != nil {
		util.APIResponse(ctx, "request invalid", 400, 0, nil)
		return
	}
	res, count, err := c.productUsecase.GetList(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "List Success", http.StatusOK, count, res)
	}
}
