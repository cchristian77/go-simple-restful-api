package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"go-simple-restful-api/app"
	"go-simple-restful-api/controller"
	"go-simple-restful-api/helper"
	"go-simple-restful-api/middleware"
	"go-simple-restful-api/repository"
	"go-simple-restful-api/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// create router
	router := app.NewRouter(categoryController)

	// create server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
