package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/koloo91/controller"
	"log"
	"os"
)

var (
	dbHost     = getEnvOrDefault("DB_HOST", "localhost")
	dbUser     = getEnvOrDefault("DB_USER", "postgres")
	dbPassword = getEnvOrDefault("DB_PASSWORD", "")
	dbName     = getEnvOrDefault("DB_NAME", "postgres")

	appUser         = getEnvOrDefault("APP_USER", "kolo")
	appUserPassword = getEnvOrDefault("APP_USER_PASSWORD", "Pass00")
)

func main() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, 5432, dbName)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("unable to open database connection: %s", err.Error())
	}

	if err := db.DB().Ping(); err != nil {
		log.Fatalf("unable to ping database: %s", err.Error())
	}

	defer db.Close()

	driver, _ := postgres.WithInstance(db.DB(), &postgres.Config{})
	migrator, _ := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err := migrator.Up(); err != nil {
		log.Println(err.Error())
	}

	db.LogMode(true)

	router := gin.Default()

	router.NoRoute(func(ctx *gin.Context) {
		ctx.File("./assets/index.html")
	})

	authorized := router.Group("", gin.BasicAuth(gin.Accounts{
		appUser: appUserPassword,
	}))

	authorized.GET("/api/ping", controller.Ping())

	{
		accounts := authorized.Group("/api/accounts")
		accounts.POST("", controller.CreateAccount(db))
		accounts.GET("", controller.GetAccounts(db))
	}

	{
		categories := authorized.Group("/api/categories")
		categories.POST("", controller.CreateCategory(db))
		categories.PUT("/:id", controller.UpdateCategory(db))
		categories.GET("", controller.GetCategories(db))
	}

	{
		bookings := authorized.Group("/api/bookings")
		bookings.POST("", controller.CreateBooking(db))
		bookings.PUT("/:id", controller.UpdateBooking(db))
		bookings.DELETE("/:id", controller.DeleteBooking(db))
		bookings.GET("", controller.GetBookings(db))
	}

	{
		balances := authorized.Group("/api/balances")
		balances.GET("", controller.GetBalances(db))
	}

	router.GET("/api/alive", controller.Alive())

	router.Static("/app", "./assets")

	log.Fatal(router.Run())
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
