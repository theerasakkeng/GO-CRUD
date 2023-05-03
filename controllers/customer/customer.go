package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theerasakkeng/GO-CRUD/cmd"
	"github.com/theerasakkeng/GO-CRUD/models"
)

func GetCustomerList(c *gin.Context) {
	db, err := cmd.UseDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	var customers []models.Customer_Response
	result := db.Raw("select * from sales.customers").Scan(&customers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": customers,
		})
	}
}

func GetCustomerDetail(c *gin.Context) {
	customer_id := c.Query("customer_id")
	db, err := cmd.UseDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	var customers []models.Customer_Response
	result := db.Raw("select * from sales.customers where customer_id = ?", customer_id).Scan(&customers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": customers,
		})
	}
}

func PostCustomer(c *gin.Context) {
	var customers models.Customer_Response
	err := c.BindJSON(&customers)
	if err != nil {
		panic(err)
	} else {
		// c.JSON(http.StatusOK, gin.H{
		// 	"data": customers,
		// })
		db, err := cmd.UseDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			db.Exec("insert into sales.customers (first_name,last_name,phone,email,street,city,state,zip_code) VALUES (?,?,?,?,?,?,?,?)", customers.First_Name, customers.Last_Name, customers.Phone, customers.Email, customers.Street, customers.City, customers.State, customers.Zip_Code)
		}
	}

	// db, err := cmd.UseDB()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, nil)
	// }
	// var customers []models.Customer_Response
	// result := db.Raw("select * from sales.customers where customer_id = ?").Scan(&customers)
	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"data": nil,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": customers,
	// 	})
	// }
}
