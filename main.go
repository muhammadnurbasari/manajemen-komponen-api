package main

import (
	"manajemen-komponen-api/app/registry"
	"manajemen-komponen-api/docs"
)

// @title Manajemen Komponen API
// @version 0.0.1
// @description This is a Manajemen Komponen server API Documentation.
// @termsOfService http://swagger.io/terms/

// @contact.name Muhammad nur basari
// @contact.email m.nurbasari@gmail.com

// @host localhost:8080
// @BasePath
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	appRegistry := registry.NewAppRegistry()
	appRegistry.StartServer()
}
