package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fabarj4/be-assignment/handler"
	"github.com/fabarj4/be-assignment/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB //created outside to make it global.
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	SecretKey string
)

func main() {

	router := gin.Default()

	// users routers
	router.GET("/users", func(ctx *gin.Context) { handler.UsersGets(ctx, DB) })
	router.POST("/register", func(ctx *gin.Context) { handler.UsersInsert(ctx, DB) })
	router.GET("/users/:id", func(ctx *gin.Context) { handler.UsersGet(ctx, DB) })
	router.GET("/users/:id/transactions_withdraw", func(ctx *gin.Context) { handler.UsersTransactionsWithdrawGets(ctx, DB) })
	router.GET("/users/:id/transactions_send", func(ctx *gin.Context) { handler.UsersTransactionsSendGets(ctx, DB) })
	router.GET("/users/:id/paymentaccounts", func(ctx *gin.Context) { handler.UsersPaymentAccountsGets(ctx, DB) })
	router.DELETE("/users/:id", func(ctx *gin.Context) { handler.UsersDelete(ctx, DB) })
	router.PUT("/users/:id", func(ctx *gin.Context) { handler.UsersUpdate(ctx, DB) })
	router.POST("/login", func(ctx *gin.Context) { handler.UsersLogin(ctx, DB) })

	// transactions routers
	router.GET("/transactions_send", func(ctx *gin.Context) { handler.TransactionsSendGets(ctx, DB) })
	router.POST("/transactions_send", func(ctx *gin.Context) { handler.TransactionsSend(ctx, DB) })
	router.GET("/transactions_send/:id", func(ctx *gin.Context) { handler.TransactionsSendGet(ctx, DB) })

	router.GET("/transactions_withdraw", func(ctx *gin.Context) { handler.TransactionsWithdrawGets(ctx, DB) })
	router.POST("/transactions_withdraw", func(ctx *gin.Context) { handler.TransactionsWithdraw(ctx, DB) })
	router.GET("/transactions_withdraw/:id", func(ctx *gin.Context) { handler.TransactionsWithdrawGet(ctx, DB) })

	// paymentaccounts routers
	router.GET("/paymentaccounts", func(ctx *gin.Context) { handler.PaymentAccountsGets(ctx, DB) })
	router.POST("/paymentaccounts", func(ctx *gin.Context) { handler.PaymentAccountsInsert(ctx, DB) })
	router.GET("/paymentaccounts/:id", func(ctx *gin.Context) { handler.PaymentAccountsGet(ctx, DB) })
	router.DELETE("/paymentaccounts/:id", func(ctx *gin.Context) { handler.PaymentAccountsDelete(ctx, DB) })
	router.PUT("/paymentaccounts/:id", func(ctx *gin.Context) { handler.PaymentAccountsUpdate(ctx, DB) })

	router.Run(":8080")
}

func init() {
	// env set
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env")
	}
	DBHost = os.Getenv("DBHost")
	DBPort = os.Getenv("DBPort")
	DBUser = os.Getenv("DBUser")
	DBName = os.Getenv("DBName")
	DBPass = os.Getenv("DBPass")

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DBHost, DBUser, DBPass, DBName, DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}
	// Auto migrate the User model
	DB.AutoMigrate(&model.Users{})
	DB.AutoMigrate(&model.PaymentAccounts{})
	DB.AutoMigrate(&model.TransactionsWithdraw{})
	DB.AutoMigrate(&model.TransactionsSend{})
}
