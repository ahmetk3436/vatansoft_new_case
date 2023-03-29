package main

import (
	"vatansoft/pkg/handler"
	"vatansoft/pkg/router"
)

func main() {
	startServer()
}

func startServer() {

	stockHandlerApi := handler.NewStockHandler()

	e := router.NewRouter(*stockHandlerApi).InitRouter()
	e.Logger.Fatal(e.Start(":1323"))
}
