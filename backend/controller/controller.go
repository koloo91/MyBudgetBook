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

	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "GIN_MODE"))

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

// Alive godoc
// @Summary Checks if the service is running
// @Description Checks if the service is running
// @Tags Alive
// @Success 204
// @Router /api/ping [get]
func ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusNoContent, "")
	}
}

// Ping godoc
// @Summary Checks if the user is logged in
// @Description Checks if the user is logged in
// @Tags Ping
// @Success 204
// @Router /api/alive [get]
func alive() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusNoContent, "")
	}
}

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account body model.AccountVo true "Create account"
// @Success 201 {object} model.AccountVo
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

// GetAccounts godoc
// @Summary Get all accounts
// @Description Get all accounts
// @Tags Accounts
// @Produce json
// @Success 200 {object} model.AccountsVo
// @Failure 400 {object} model.ErrorVo
// @Router /api/accounts [get]
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

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body model.CategoryVo true "Create category"
// @Success 201 {object} model.CategoryVo
// @Failure 400 {object} model.ErrorVo
// @Router /api/categories [post]
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

// UpdateCategory godoc
// @Summary Updates a category
// @Description Updates a category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body model.CategoryVo true "Update category"
// @Param id path string true "Category id"
// @Success 200 {object} model.CategoryVo
// @Failure 400 {object} model.ErrorVo
// @Router /api/categories/{id} [put]
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

// GetCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags Categories
// @Produce json
// @Success 200 {object} model.CategoriesVo
// @Failure 400 {object} model.ErrorVo
// @Router /api/categories [get]
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

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new booking
// @Tags Bookings
// @Accept json
// @Produce json
// @Param category body model.BookingVo true "Create booking"
// @Success 201 {object} model.BookingVo
// @Failure 400 {object} model.ErrorVo
// @Router /api/bookings [post]
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

// UpdateBooking godoc
// @Summary Updates a booking
// @Description Updates a booking
// @Tags Bookings
// @Accept json
// @Produce json
// @Param category body model.BookingVo true "Create booking"
// @Param id path string true "Booking id"
// @Param updateStrategy query string false "update only this entry or all of the standing order" Enums(ONE, ALL) default(ONE)
// @Success 200 {object} model.BookingVo
// @Failure 400 {object} model.ErrorVo
// @Router /api/bookings/{id} [put]
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

// GetBookings godoc
// @Summary Get bookings in the given time range
// @Description Get bookings in the given time range
// @Tags Bookings
// @Produce json
// @Param startDate query string false "start date of the range: 2006-01-02T15:04:05Z07:00"
// @Param endDate query string false "end date of the range: 2006-01-02T15:04:05Z07:00"
// @Success 200 {object} model.BookingsVo
// @Failure 400 {object} model.ErrorVo
// @Router /api/bookings [get]
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

// DeleteBooking godoc
// @Summary Deletes a booking
// @Description Deletes a booking
// @Tags Bookings
// @Param id path string true "Booking id"
// @Success 204
// @Failure 400 {object} model.ErrorVo
// @Router /api/bookings/{id} [delete]
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

// GetBalances godoc
// @Summary Get current balances
// @Description Get current balances
// @Tags Balances
// @Produce json
// @Success 200 {object} model.AccountBalancesVo
// @Failure 400 {object} model.ErrorVo
// @Router /api/categories [get]
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
