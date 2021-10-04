package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email,omitempty" gorm:"uniqueIndex"`
	Phone     string `json:"phone,omitempty"`
	/*CompanyID *uint    `json:"company_id,omitempty"`
	Company   *Company `json:"company,omitempty"`*/
}

/*type Company struct {
	gorm.Model
	Id int `json:"id" gorm:"uniqueIndex"`
	Name         string          `json:"name" gorm:"uniqueIndex"`
	Description  string          `json:"description,omitempty"`
	CategoryID   *uint            `json:"category_id,omitempty"`
	Category     *CompanyCategory `json:"category"`
	Website      string          `json:"website,omitempty"`
}

type CompanyCategory struct {
	gorm.Model
	Id int `json:"id" gorm:"uniqueIndex"`
	Name string `json:"name" gorm:"uniqueIndex"`
}

type Response struct {
	Code    int         `json:"code,omitempty"`
	Body    interface{} `json:"body,omitempty"`
	Title   string      `json:"title,omitempty"`
	Message string      `json:"message,omitempty"`
}*/
