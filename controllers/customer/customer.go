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
