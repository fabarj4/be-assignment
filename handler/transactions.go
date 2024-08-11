package handler

import (
	"net/http"
	"strings"

	"github.com/fabarj4/be-assignment/lib"
	"github.com/fabarj4/be-assignment/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionsSend(c *gin.Context, db *gorm.DB) {
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
	data := model.TransactionsSend{}
	if err := c.BindJSON(&data); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// process transaction
	data.Process(5)

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

func TransactionsSendGet(c *gin.Context, db *gorm.DB) {
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
	data := &model.TransactionsSend{}
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

func TransactionsSendGets(c *gin.Context, db *gorm.DB) {
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
	Transactions := []model.TransactionsSend{}
	if err := db.Find(&Transactions).Error; err != nil {
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
		Data:    Transactions,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func TransactionsWithdraw(c *gin.Context, db *gorm.DB) {
	data := model.TransactionsWithdraw{}
	if err := c.BindJSON(&data); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// process transaction
	data.Process(5)

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

func TransactionsWithdrawGet(c *gin.Context, db *gorm.DB) {
	data := &model.TransactionsWithdraw{}
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

func TransactionsWithdrawGets(c *gin.Context, db *gorm.DB) {
	Transactions := []model.TransactionsWithdraw{}
	if err := db.Find(&Transactions).Error; err != nil {
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
		Data:    Transactions,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}
