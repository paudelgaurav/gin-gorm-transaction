package main

import (
	"github.com/paudelgaurav/gin-gorm-transaction/model"
	"github.com/paudelgaurav/gin-gorm-transaction/route"
)

func main() {

	db, _ := model.DBConnection()
	route.SetUpRoute(db)
}
