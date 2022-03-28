package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//StatusInList -> checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return true
}

// DBTransactionMiddleware : to setup the database transaction middleware.
func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		txHandle := db.Begin()
		log.Print("Beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		// Set is used to store a new key/value pair exclusively for this context.
		ctx.Set("db_trx", txHandle)

		// Next should be used only inside middleware.
		// It executes the pending handlers in the chain inside the calling handler.
		ctx.Next()

		if StatusInList(ctx.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			log.Print("commiting transaction")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("trx commit error ", err)
			}
		} else {
			log.Print("Rolling back transaction due to invalid status code: ", ctx.Writer.Status())
			txHandle.Rollback()
		}
	}

}
