package router

import (
	"vatansoft/pkg/handler"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Handler handler.Handler
}

func NewRouter(handler handler.Handler) *Router {
	return &Router{
		Handler: handler,
	}
}

func (r *Router) InitRouter() *echo.Echo {
	e := echo.New()
	e.GET("/stocks", r.Handler.Api.GetAllProductApi)
	e.GET("/stocks/filter", r.Handler.Api.FilterSearchStockProductApi)
	e.POST("/stock/insert", r.Handler.Api.CreateStockProductApi)
	e.PUT("/stock/:id/update", r.Handler.Api.UpdateStockProductApi)
	e.DELETE("/stock/:id/delete", r.Handler.Api.DeleteStockProductApi)
	e.GET("/stock/:id", r.Handler.Api.GetStockProductByIdApi)

	e.POST("/stock/category/insert", r.Handler.Api.InsertCategoryForAllProductApi)
	e.DELETE("/stock/:id/category/:id/delete", r.Handler.Api.DeleteCategoryForProductByIdApi)
	e.DELETE("/stock/:id/category/delete", r.Handler.Api.DeleteCategoryForProductsByIdApi)

	e.GET("/categories", r.Handler.CategoryApi.GetAllCategoriesApi)
	e.GET("/category/:id", r.Handler.CategoryApi.GetCategoryByIdApi)
	e.POST("/category/insert", r.Handler.CategoryApi.CreateCategoryApi)
	e.DELETE("/category/:id/delete", r.Handler.CategoryApi.DeleteCategoryApi)
	e.PUT("/category/:id/update", r.Handler.CategoryApi.UpdateCategoryApi)

	//bill crud
	e.GET("/bills", r.Handler.BillApi.GetAllBillsApi)
	e.GET("/bill/:id", r.Handler.BillApi.GetBillByIdApi)
	e.POST("/bill/insert", r.Handler.BillApi.CreateBillApi)
	e.DELETE("/bill/:id/delete", r.Handler.BillApi.DeleteBillApi)
	e.PUT("/bill/:id/update", r.Handler.BillApi.UpdateBillApi)
	//property crud
	e.GET("/propertys", r.Handler.PropertyApi.GetAllPropertysApi)
	e.GET("/property/:id", r.Handler.PropertyApi.GetPropertyByIdApi)
	e.POST("/property/insert", r.Handler.PropertyApi.CreatePropertyApi)
	e.DELETE("/property/:id/delete", r.Handler.PropertyApi.DeletePropertyApi)
	e.PUT("/property/:id/update", r.Handler.PropertyApi.UpdatePropertyApi)
	return e
}
