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

type Customer_Req struct {
	First_Name string  `json:"first_name"`
	Last_Name  string  `json:"last_name"`
	Phone      *string `json:"phone"`
	Email      string  `json:"email"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Zip_Code   string  `json:"zip_code"`
}

func (Customer_Req) TableName() string {
	return "sales.customers"
}
