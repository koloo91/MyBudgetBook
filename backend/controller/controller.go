package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/koloo91/mapper"
	"github.com/koloo91/model"
	"github.com/koloo91/service"
	"log"
	"net/http"
	"time"
)

func SetupRoutes(db *gorm.DB, appUser, appUserPassword string) *gin.Engine {
	router := gin.Default()

	authorized := router.Group("", gin.BasicAuth(gin.Accounts{
		appUser: appUserPassword,
	}))

	authorized.GET("/api/ping", ping())

	{
		accounts := authorized.Group("/api/accounts")
		accounts.POST("", createAccount(db.DB()))
		accounts.GET("", getAccounts(db.DB()))
	}

	{
		categories := authorized.Group("/api/categories")
		categories.POST("", createCategory(db.DB()))
		categories.PUT("/:id", updateCategory(db.DB()))
		categories.GET("", getCategories(db.DB()))
	}

	{
		bookings := authorized.Group("/api/bookings")
		bookings.POST("", createBooking(db.DB()))
		bookings.PUT("/:id", updateBooking(db.DB()))
		bookings.DELETE("/:id", deleteBooking(db.DB()))
		bookings.GET("", getBookings(db.DB()))
	}

	{
		balances := authorized.Group("/api/balances")
		balances.GET("", getBalances(db.DB()))
	}

	router.GET("/api/alive", alive())

	return router
}

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

func createAccount(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accountVo model.AccountVo
		if err := ctx.ShouldBindJSON(&accountVo); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		createdAccount, err := service.CreateAccount(ctx.Request.Context(), db, mapper.AccountVoToEntity(accountVo))
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
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
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
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
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		createdCategory, err := service.CreateCategory(ctx.Request.Context(), db, mapper.CategoryVoToEntity(categoryVo))
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
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
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		updatedCategory, err := service.UpdateCategory(ctx.Request.Context(), db, id, mapper.CategoryVoToEntity(categoryVo))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, mapper.CategoryEntityToVo(updatedCategory))
	}
}

func getCategories(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		categories, err := service.GetCategories(ctx.Request.Context(), db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, model.CategoriesVo{Content: mapper.CategoryEntitiesToVos(categories)})
	}
}

func createBooking(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookingVo model.BookingVo
		if err := ctx.ShouldBindJSON(&bookingVo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		createdBooking, err := service.CreateBooking(ctx.Request.Context(), db, mapper.BookingVoToEntity(bookingVo))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
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
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		updatedBooking, err := service.UpdateBooking(ctx.Request.Context(), db, id, mapper.BookingVoToEntity(bookingVo), updateStrategy)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
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
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		endDate, err := time.Parse(time.RFC3339, endDateString)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		bookings, err := service.GetBookings(ctx.Request.Context(), db, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
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
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, "")
	}
}

func getBalances(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		balances, err := service.GetBalances(ctx.Request.Context(), db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, model.AccountBalancesVo{Content: mapper.AccountBalanceEntitiesToVos(balances)})
	}
}
