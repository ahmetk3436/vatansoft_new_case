package router

import (
	"vatansoft/pkg/handler"

	"github.com/labstack/echo/v4"
)

type Router struct {
	StockHandler handler.StockHandler
}

func NewRouter(stockHandler handler.StockHandler) *Router {
	return &Router{
		StockHandler: stockHandler,
	}
}

func (r *Router) InitRouter() *echo.Echo {
	e := echo.New()
	e.PUT("/stock/insert", r.StockHandler.Api.CreateStockProductApi)
	/* 	e.GET("/stocks", r.StockHandler.Api.CreateStockProductService)
	   	e.GET("/stocks/filter", r.StockHandler.Api.FilterProductsHandler)
	   	e.PUT("/stock/insert", r.StockHandler.Api.CreateStockProductService)
	   	e.POST("/stock/:id/update", r.StockHandler.Api.UpdateProductHandler)
	   	e.DELETE("/stock/:id/delete", r.StockHandler.Api.DeleteProductHandler)
	   	e.GET("/stock/:id", r.StockHandler.Api.GetProductHandler)

	   	e.POST("/stock/category/insert", r.StockHandler.Api.CreateCategoryHandler)
	   	e.DELETE("/stock/:id/category/:id/delete", r.StockHandler.Api.DeleteProductCategoryHandler)
	   	e.DELETE("/stock/:id/category/delete", r.StockHandler.Api.DeleteAllProductCategoriesHandler)

	   	e.GET("/categories", r.StockHandler.Api.GetAllCategoriesHandler)
	   	e.GET("/category/:id", r.StockHandler.Api.GetCategoryHandler)
	   	e.POST("/category/insert", r.StockHandler.Api.CreateCategoryHandler)
	   	e.DELETE("/category/:id/delete", r.StockHandler.Api.DeleteCategoryHandler)
	   	e.PUT("/category/:id/update", r.StockHandler.Api.UpdateCategoryHandler) */

	return e
}
