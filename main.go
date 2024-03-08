package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shkuran/go-library-microservices/book-service/book"
	"github.com/shkuran/go-library-microservices/book-service/config"
	"github.com/shkuran/go-library-microservices/book-service/db"
	"github.com/shkuran/go-library-microservices/book-service/routes"
)

func main() {
	conf := config.LoadConfig()
	port := conf.Database.Port
	user := conf.Database.User
	pass := conf.Database.Password
	dbName := conf.Database.DbName
	sslMode := conf.Database.SslMode
	driverName := conf.Database.DriverName
	connStr := fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=%s",
		port, user, pass, dbName, sslMode)

	varDb, err := db.InitDB(driverName, connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	//db.CreateTables(varDb)

	server := gin.Default()

	bookRepo := book.NewRepo(varDb)
	bookHandler := book.NewHandler(bookRepo)

	routes.RegisterRoutes(server, bookHandler)

	err = server.Run(":" + conf.Server.Port)
	if err != nil {
		log.Fatal(err)
		return
	}
}
