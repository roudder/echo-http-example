package handler

import (
	"awesomeProject/pkg/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DictionaryHandler struct {
	ds service.DictService
}

type DictHandler interface {
	Get() string
	Set(key string, value string)
}

func (dh *DictionaryHandler) Get(c echo.Context) error {
	key := c.Param("key")
	//just for examp	le, theoretically key can be ""
	value := dh.ds.Get(key)
	if value == "" {
		c.String(http.StatusNotFound, "no such key")
	}
	return c.String(http.StatusOK, value)
}

func (dh *DictionaryHandler) Set(c echo.Context) error {
	key := c.FormValue("key")
	value := c.FormValue("value")
	dh.ds.Set(key, value)
	return c.String(http.StatusOK, fmt.Sprintf("key: %s, value:%s  settled", key, value))

}

func NewDictionaryHandler(ds service.DictService) *DictionaryHandler {
	return &DictionaryHandler{
		ds: ds,
	}
}
