package main

import (
	"log"
	auth "perpustakaan/Auth"
	handler "perpustakaan/Handler"
	user "perpustakaan/User"
	siswa "perpustakaan/User/Siswa"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_perpustakaan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	siswaRepository := siswa.NewRepository(db)

	userService := user.NewService(userRepository)
	siswaService := siswa.NewService(siswaRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	siswaHandler := handler.NewSiswaHandler(siswaService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/siswa", siswaHandler.RegisterSiswa)
	api.POST("/user", userHandler.RegisterUser)
	router.Run()
}
