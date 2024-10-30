package main

import (
	"exclusive-web/web/middleware"
	"exclusive-web/web/modules/category"
	"exclusive-web/web/modules/product"
	"exclusive-web/web/modules/user"
	db "exclusive-web/web/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	var (
		err error
	)

	router.Use(
		middleware.AllowCORS(),
	)

	loadEnv := godotenv.Load()

	if loadEnv != nil {
		log.Fatal("Error loading .env file")
		return
	}

	router.Static("/static", "./assets/images")

	dsn := db.GormPostgres("host=localhost user=postgres password=Lumbanpaung,050490 dbname=exclusive port=5432 sslmode=disable TimeZone=Asia/Jakarta")

	userHandler := user.NewRequestHandler(dsn)
	userHandler.HandlerUser(router)

	categoryHandler := category.NewRequestHandler(dsn)
	categoryHandler.HandlerRepository(router)

	productHandler := product.NewRequestHandler(dsn)
	productHandler.HandlerProduct(router)

	err = router.Run(":8888")

	if err != nil {
		log.Println("main router.Run:", err, dsn)
		return
	}
}
