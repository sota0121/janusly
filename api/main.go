package main

import (
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Gorilla Mux Routerを使用
	router := mux.NewRouter()
	e.Any("/*", echo.WrapHandler(router))

	// URL短縮エンドポイントの追加
	router.HandleFunc("/shorten", ShortenURL).Methods("POST")
	router.HandleFunc("/{shortURL}", RedirectURL).Methods("GET")

	e.Logger.Fatal(e.Start(":8080"))
}
