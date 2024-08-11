package handler

import (
	"net/http"
	"strings"

	"github.com/fabarj4/be-assignment/lib"
	"github.com/fabarj4/be-assignment/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentAccountsInsert(c *gin.Context, db *gorm.DB) {
	// check token
	auth := c.Request.Header["Authorization"]
	if len(auth) == 0 {
		res := HandlerResponse{
			Message: "token not found",
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// validate token
	token := strings.Split(auth[0], "Bearer ")[len(strings.Split(auth[0], "Bearer"))-1]
	if err := lib.VerifyToken(token); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	data := model.PaymentAccounts{}
	if err := c.BindJSON(&data); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if err := db.Create(&data).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := HandlerResponse{
		Message: "success",
		Data:    data,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func PaymentAccountsUpdate(c *gin.Context, db *gorm.DB) {
	// check token
	auth := c.Request.Header["Authorization"]
	if len(auth) == 0 {
		res := HandlerResponse{
			Message: "token not found",
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// validate token
	token := strings.Split(auth[0], "Bearer ")[len(strings.Split(auth[0], "Bearer"))-1]
	if err := lib.VerifyToken(token); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	data := &model.PaymentAccounts{}
	id := c.Param("id")
	if err := db.First(data, id).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return

	}
	if err := c.BindJSON(data); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if err := db.Save(data).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := HandlerResponse{
		Message: "success",
		Data:    data,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func PaymentAccountsDelete(c *gin.Context, db *gorm.DB) {
	// check token
	auth := c.Request.Header["Authorization"]
	if len(auth) == 0 {
		res := HandlerResponse{
			Message: "token not found",
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// validate token
	token := strings.Split(auth[0], "Bearer ")[len(strings.Split(auth[0], "Bearer"))-1]
	if err := lib.VerifyToken(token); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	data := &model.PaymentAccounts{}
	id := c.Param("id")
	// check data first
	if err := db.First(data, id).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if err := db.Delete(data, id).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := HandlerResponse{
		Message: "success",
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func PaymentAccountsGet(c *gin.Context, db *gorm.DB) {
	// check token
	auth := c.Request.Header["Authorization"]
	if len(auth) == 0 {
		res := HandlerResponse{
			Message: "token not found",
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// validate token
	token := strings.Split(auth[0], "Bearer ")[len(strings.Split(auth[0], "Bearer"))-1]
	if err := lib.VerifyToken(token); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	data := &model.PaymentAccounts{}
	id := c.Param("id")
	if err := db.First(data, id).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return

	}
	res := HandlerResponse{
		Message: "success",
		Data:    data,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func PaymentAccountsGets(c *gin.Context, db *gorm.DB) {
	// check token
	auth := c.Request.Header["Authorization"]
	if len(auth) == 0 {
		res := HandlerResponse{
			Message: "token not found",
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// validate token
	token := strings.Split(auth[0], "Bearer ")[len(strings.Split(auth[0], "Bearer"))-1]
	if err := lib.VerifyToken(token); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	PaymentAccounts := []model.PaymentAccounts{}
	if err := db.Find(&PaymentAccounts).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := HandlerResponse{
		Message: "success",
		Data:    PaymentAccounts,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}
