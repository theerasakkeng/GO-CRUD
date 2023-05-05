package models

type Customer_Response struct {
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

type Customer_Request struct {
	First_Name string  `json:"first_name" binding:"required"`
	Last_Name  string  `json:"last_name" binding:"required"`
	Phone      *string `json:"phone"`
	Email      string  `json:"email" binding:"required"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Zip_Code   string  `json:"zip_code"`
}

type Customer_Id struct {
	Customer_Id int `json:"customer_id"`
}

type Tabler interface {
	TableName() string
}

func (Customer_Response) TableName() string {
	return "sales.customers"
}

func (Customer_Request) TableName() string {
	return "sales.customers"
}
