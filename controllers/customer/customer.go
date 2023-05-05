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
	sqlDB, err := db.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	defer sqlDB.Close()
	var customers []models.Customer_Response
	result := db.Find(&customers)
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
	sqlDB, err := db.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	defer sqlDB.Close()
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
	var customers models.Customer_Request
	err := c.BindJSON(&customers)
	if err != nil {
		panic(err)
	} else {
		db, err := cmd.UseDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			result := db.Create(&customers)
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
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
		defer sqlDB.Close()
	}
}

func UpdateCustomer(c *gin.Context) {
	var customers models.Customer_Request
	err := c.BindJSON(&customers)
	if err != nil {
		panic(err)
	} else {
		db, err := cmd.UseDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			result := db.Where("customer_id = ?", 2446).Save(&customers)
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
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
		defer sqlDB.Close()
	}
}

func DeleteCustomer(c *gin.Context) {
	var customers models.Customer_Request
	var customer_id models.Customer_Id
	err := c.BindJSON(&customer_id)
	if err != nil {
		panic(err)
	} else {
		db, err := cmd.UseDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			result := db.Where("customer_id = ?", customer_id.Customer_Id).Delete(&customers)
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
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
		defer sqlDB.Close()
	}
}
