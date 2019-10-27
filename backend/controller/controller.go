package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/koloo91/mapper"
	"github.com/koloo91/model"
	"github.com/koloo91/service"
	"net/http"
)

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusNoContent, "")
	}
}

func CreateAccount(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accountVo model.AccountVo
		if err := ctx.ShouldBindJSON(&accountVo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		createdAccount, err := service.CreateAccount(db, mapper.AccountVoToEntity(accountVo))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, mapper.AccountEntityToVo(createdAccount))
	}
}

func GetAccounts(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accounts, err := service.GetAccounts(db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, model.AccountsVo{Content: mapper.AccountEntitiesToVos(accounts)})
	}
}

func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var categoryVo model.CategoryVo
		if err := ctx.ShouldBindJSON(&categoryVo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		createdCategory, err := service.CreateCategory(db, mapper.CategoryVoToEntity(categoryVo))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, mapper.CategoryEntityToVo(createdCategory))
	}
}

func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var categoryVo model.CategoryVo
		if err := ctx.ShouldBindJSON(&categoryVo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		createdCategory, err := service.UpdateCategory(db, id, mapper.CategoryVoToEntity(categoryVo))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, mapper.CategoryEntityToVo(createdCategory))
	}
}

func GetCategories(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		categories, err := service.GetCategories(db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, model.CategoriesVo{Content: mapper.CategoryEntitiesToVos(categories)})
	}
}

func CreateBooking(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookingVo model.BookingVo
		if err := ctx.ShouldBindJSON(&bookingVo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		createdBooking, err := service.CreateBooking(db, mapper.BookingVoToEntity(bookingVo))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, mapper.BookingEntityToVo(createdBooking))
	}
}

func GetBookings(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bookings, err := service.GetBookings(db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorVo{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, model.BookingsVo{Content: mapper.BookingEntitiesToVos(bookings)})
	}
}
