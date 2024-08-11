package handler

import (
	"net/http"
	"strings"

	"github.com/fabarj4/be-assignment/lib"
	"github.com/fabarj4/be-assignment/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsersInsert(c *gin.Context, db *gorm.DB) {
	user := model.Users{}
	if err := c.BindJSON(&user); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	hashPassword, err := lib.HashPassword(user.Password)
	if err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return

	}
	user.Password = hashPassword
	if err := db.Create(&user).Error; err != nil {
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
		Data:    user,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func UsersUpdate(c *gin.Context, db *gorm.DB) {
	user := &model.Users{}
	id := c.Param("id")
	if err := db.First(user, id).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return

	}
	if err := c.BindJSON(user); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if err := db.Save(user).Error; err != nil {
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
		Data:    user,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func UsersDelete(c *gin.Context, db *gorm.DB) {
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
	user := &model.Users{}
	id := c.Param("id")
	// check data first
	if err := db.First(user, id).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if err := db.Delete(user, id).Error; err != nil {
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

func UsersGet(c *gin.Context, db *gorm.DB) {
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
	user := &model.Users{}
	id := c.Param("id")
	if err := db.First(user, id).Error; err != nil {
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
		Data:    user,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func UsersGets(c *gin.Context, db *gorm.DB) {
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

	users := []model.Users{}
	if err := db.Find(&users).Error; err != nil {
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
		Data:    users,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}
func UsersLogin(c *gin.Context, db *gorm.DB) {
	user := model.Users{}
	if err := c.BindJSON(&user); err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	userDB := &model.Users{}
	// check data first
	if err := db.Where("email = ?", user.Email).First(&userDB).Error; err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if valid := lib.CheckPasswordHash(user.Password, userDB.Password); !valid {
		res := HandlerResponse{
			Message: "email/password wrong",
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	token, err := lib.CreateToken(int(userDB.ID))
	if err != nil {
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
		Data:    token,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)

}

func UsersTransactionsWithdrawGets(c *gin.Context, db *gorm.DB) {
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
	id := c.Param("id")
	data := []*model.UsersTransactionsWithdraws{}

	query := `
	SELECT
	transactions_withdraws.id,payment_accounts_id,transactions_withdraws.users_id,currency,amount,
	discount_code,amount_after,status,to_address,
	users.name as users_name,
	payment_accounts.account_name as payment_accounts_name
	FROM transactions_withdraws
	INNER JOIN users ON users.id = transactions_withdraws.users_id
	INNER JOIN payment_accounts ON payment_accounts.id = transactions_withdraws.payment_accounts_id
	WHERE transactions_withdraws.users_id = ? 
	`

	rows, err := db.Raw(query, id).Rows()
	if err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	defer rows.Close()
	for rows.Next() {
		temp := &model.UsersTransactionsWithdraws{}
		if err := db.ScanRows(rows, &temp); err != nil {
			res := HandlerResponse{
				Message: err.Error(),
				Data:    nil,
				Status:  http.StatusInternalServerError,
			}
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		data = append(data, temp)
	}

	res := HandlerResponse{
		Message: "success",
		Data:    data,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func UsersTransactionsSendGets(c *gin.Context, db *gorm.DB) {
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
	id := c.Param("id")
	data := []*model.UsersTransactionsSends{}

	query := `
	SELECT
	transactions_sends.id,
	transactions_sends.users_id,
	currency,amount,status,to_address,
	(select name from users where users.id = transactions_sends.users_id) as users_name,
	(select name from users where users.id = transactions_sends.to_users_id) as to_users_name
	FROM transactions_sends
	WHERE transactions_sends.users_id = ? 
	`

	rows, err := db.Raw(query, id).Rows()
	if err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	defer rows.Close()
	for rows.Next() {
		temp := &model.UsersTransactionsSends{}
		if err := db.ScanRows(rows, &temp); err != nil {
			res := HandlerResponse{
				Message: err.Error(),
				Data:    nil,
				Status:  http.StatusInternalServerError,
			}
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		data = append(data, temp)
	}

	res := HandlerResponse{
		Message: "success",
		Data:    data,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}

func UsersPaymentAccountsGets(c *gin.Context, db *gorm.DB) {
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
	id := c.Param("id")
	data := []*model.UsersPaymentAccounts{}

	query := `
	SELECT
	payment_accounts.id,account_name,account_number,active,
	users.id as users_id,users.name as users_name
	FROM payment_accounts
	INNER JOIN users ON users.id = payment_accounts.users_id
	WHERE users.id = ? 
	`

	rows, err := db.Raw(query, id).Rows()
	if err != nil {
		res := HandlerResponse{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	defer rows.Close()
	for rows.Next() {
		temp := &model.UsersPaymentAccounts{}
		if err := db.ScanRows(rows, &temp); err != nil {
			res := HandlerResponse{
				Message: err.Error(),
				Data:    nil,
				Status:  http.StatusInternalServerError,
			}
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		data = append(data, temp)
	}

	res := HandlerResponse{
		Message: "success",
		Data:    data,
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, res)
}
