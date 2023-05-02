package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Customer_Model struct {
	Customer_Id int     `json:"customer_id"`
	First_Name  string  `json:"first_name"`
	Last_Name   string  `json:"last_name"`
	Phone       *string `json:"phone"`
	Email       string  `json:"email"`
	Street      string  `json:"street"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Zip_Code    string  `json:"zip_code"`
}

var db *gorm.DB

func main() {
	var err error
	dsn := "sqlserver://sa:Keng1234@localhost?database=TestDB"
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var customers []Customer_Model
		result := db.Raw("select * from sales.customers where state = ? OR state = ?", "NY", "CA").Scan(&customers)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": customers,
			})
		}
	})
	r.Run()

}
