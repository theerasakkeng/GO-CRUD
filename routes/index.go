package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theerasakkeng/GO-CRUD/controllers/customer"
)

var Router = gin.Default()

func InitRoute() {
	CustomerRoute := Router.Group("/api/customer")
	{
		CustomerRoute.GET("/getcustomers", customer.GetCustomerList)
		CustomerRoute.GET("/getcustomerdetail", customer.GetCustomerDetail)

	}
}
