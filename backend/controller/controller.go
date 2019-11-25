package controller

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/koloo91/docs"
	"github.com/koloo91/mapper"
	"github.com/koloo91/model"
	"github.com/koloo91/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"time"
)

// gin-swagger middleware
// swagger embed files

func SetupRoutes(db *sql.DB, appUser, appUserPassword string) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(unhandledErrorHandler())

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	authorized := router.Group("", gin.BasicAuth(gin.Accounts{
		appUser: appUserPassword,
	}))

	authorized.GET("/api/ping", ping())

	{
		accounts := authorized.Group("/api/accounts")
		accounts.POST("", createAccount(db))
		accounts.GET("", getAccounts(db))
	}

	{
		categories := authorized.Group("/api/categories")
		categories.POST("", createCategory(db))
		categories.PUT("/:id", updateCategory(db))
		categories.GET("", getCategories(db))
	}

	{
		bookings := authorized.Group("/api/bookings")
		bookings.POST("", createBooking(db))
		bookings.PUT("/:id", updateBooking(db))
		bookings.DELETE("/:id", deleteBooking(db))
		bookings.GET("", getBookings(db))
	}

	{
		balances := authorized.Group("/api/balances")
		balances.GET("", getBalances(db))
	}

	router.GET("/api/alive", alive())

	url := ginSwagger.URL("router/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func unhandledErrorHandler() gin.HandlerFunc {
	gin.Default()
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorVo{Message: "unexpected error"})
			}
		}()
		ctx.Next()
	}
}

// Ping godoc
// @Summary Checks if the user is logged in
// @Description Checks if the user is logged in
// @Tags Ping
// @Success 204
// @Router /api/ping [get]
func ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusNoContent, "")
	}
}

func alive() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusNoContent, "")
	}
}

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new account
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param account body model.AccountVo true "Create account"
// @Success 200 {object} model.AccountVo
// @Failure 400 {object} model.ErrorVo
// @Router /api/accounts [post]
func createAccount(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accountVo model.AccountVo
		if err := ctx.ShouldBindJSON(&accountVo); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		createdAccount, err := service.CreateAccount(ctx.Request.Context(), db, mapper.AccountVoToEntity(accountVo))
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, mapper.AccountEntityToVo(createdAccount))
	}
}

func getAccounts(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accounts, err := service.GetAccounts(ctx.Request.Context(), db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, model.AccountsVo{Content: mapper.AccountEntitiesToVos(accounts)})
	}
}

func createCategory(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var categoryVo model.CategoryVo
		if err := ctx.ShouldBindJSON(&categoryVo); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		createdCategory, err := service.CreateCategory(ctx.Request.Context(), db, mapper.CategoryVoToEntity(categoryVo))
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, mapper.CategoryEntityToVo(createdCategory))
	}
}

func updateCategory(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var categoryVo model.CategoryVo
		if err := ctx.ShouldBindJSON(&categoryVo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		updatedCategory, err := service.UpdateCategory(ctx.Request.Context(), db, id, mapper.CategoryVoToEntity(categoryVo))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, mapper.CategoryEntityToVo(updatedCategory))
	}
}

func getCategories(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		categories, err := service.GetCategories(ctx.Request.Context(), db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, model.CategoriesVo{Content: mapper.CategoryEntitiesToVos(categories)})
	}
}

func createBooking(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookingVo model.BookingVo
		if err := ctx.ShouldBindJSON(&bookingVo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		createdBooking, err := service.CreateBooking(ctx.Request.Context(), db, mapper.BookingVoToEntity(bookingVo))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, mapper.BookingEntityToVo(createdBooking))
	}
}

func updateBooking(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		updateStrategy := ctx.DefaultQuery("updateStrategy", service.UpdateStrategyOne)

		var bookingVo model.BookingVo
		if err := ctx.ShouldBindJSON(&bookingVo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		updatedBooking, err := service.UpdateBooking(ctx.Request.Context(), db, id, mapper.BookingVoToEntity(bookingVo), updateStrategy)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, mapper.BookingEntityToVo(updatedBooking))
	}
}

func getBookings(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startDateString := ctx.DefaultQuery("startDate", service.BeginningOfMonth().Format(time.RFC3339))
		endDateString := ctx.DefaultQuery("endDate", service.EndOfMonth().Format(time.RFC3339))

		startDate, err := time.Parse(time.RFC3339, startDateString)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		endDate, err := time.Parse(time.RFC3339, endDateString)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		bookings, err := service.GetBookings(ctx.Request.Context(), db, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, model.BookingsVo{Content: mapper.BookingEntitiesToVos(bookings)})
	}
}

func deleteBooking(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		deleteStrategy := ctx.DefaultQuery("deleteStrategy", service.DeleteStrategyOne)

		err := service.DeleteBooking(ctx.Request.Context(), db, id, deleteStrategy)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, "")
	}
}

func getBalances(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		balances, err := service.GetBalances(ctx.Request.Context(), db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, model.AccountBalancesVo{Content: mapper.AccountBalanceEntitiesToVos(balances)})
	}
}
