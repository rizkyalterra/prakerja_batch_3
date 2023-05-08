package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `jsin:"string"`
}

type BaseResponse struct {
	Status  bool        `json:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	connectDatabase()
	e := echo.New()
	e.GET("/users", GetUserController)
	e.POST("/users", InsertUserController)
	e.Start(":8000")
}

func connectDatabase() {
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja_3?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Init database failed")
	}
	migratDatabase()
}

func migratDatabase() {
	DB.AutoMigrate(&User{})
}

func GetUserController(c echo.Context) error {
	var users []User

	result := DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, BaseResponse{
			false, "Gagal get data user dari database", nil,
		})
	}

	return c.JSON(200, BaseResponse{
		true, "success", users,
	})
}

func InsertUserController(c echo.Context) error {
	var insertUser User
	c.Bind(&insertUser)

	// masukkan ke database
	result := DB.Create(&insertUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, BaseResponse{
			false, "Gagal insert ke database", nil,
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		true, "berhasil insert database user", insertUser,
	})
}
