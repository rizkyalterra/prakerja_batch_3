package controllers

import (
	"fmt"
	"net/http"

	"prakerja3/configs"
	"prakerja3/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUserController(c echo.Context) error {
	var users []models.User

	result := configs.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal get data user dari database", Data: nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status: true, Message: "success", Data: users,
	})
}

func InsertUserController(c echo.Context) error {
	var insertUser models.User
	c.Bind(&insertUser)
	fmt.Println(insertUser)

	// logic bisnis
	// di cek database ada ?
	// 409
	result := configs.DB.First(&models.User{}, "email = ?", insertUser.Email)
	if result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusConflict, models.BaseResponse{
			Status: false, Message: "Email telah ada", Data: nil,
		})
	}

	// masukkan ke database
	result = configs.DB.Create(&insertUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal insert ke database", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "berhasil insert database user", Data: insertUser,
	})
}
