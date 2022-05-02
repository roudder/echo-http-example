package main

import (
	"awesomeProject/pkg/handler"
	"awesomeProject/pkg/repository"
	"awesomeProject/pkg/service"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// DI stuff
	dictRep := repository.NewDictionary()
	dictService := service.NewDictionaryService(dictRep)
	dictHandler := handler.NewDictionaryHandler(dictService)

	// Routes
	e.GET("/dictionary/:key", dictHandler.Get)
	e.POST("/dictionary",  dictHandler.Set)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
