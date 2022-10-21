package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	Customer struct {
		Name         string    `json:"cName"`
		Tel          uint64    `json:"cTel"`
		Address      string    `json:"cAddress"`
		ID           uint64    `json:"cID"`
		RegisterDate time.Time `json:"cRegisterDate"`
	}

	CustomerResponse struct {
		Name         string `json:"cName"`
		Tel          uint64 `json:"cTel"`
		Address      string `json:"cAddress"`
		ID           uint64 `json:"cID"`
		RegisterDate string `json:"cRegisterDate"`
		Message      string `json:"msg"`
	}

	AllCustomers struct {
		Size      uint64              `json:"size"`
		Customers []*CustomerResponse `json:"customers"`
		Message   string              `json:"msg"`
	}

	JustMSG struct {
		Message string `json:"msg"`
	}

	ReportResponse struct {
		TotalCustomers uint64 `json:"totalCustomers"`
		Period         int    `json:"period"`
		Message        string `json:"msg"`
	}
)

var (
	db          = map[uint64]*Customer{}
	last uint64 = 0
)

func CreateNewCustomer(c echo.Context) error {

	newUser := &Customer{
		ID:           last + 1,
		RegisterDate: time.Now(),
	}
	if err := c.Bind(newUser); err != nil {
		return err
	}
	db[newUser.ID] = newUser
	last++

	response := newCustomerResponse(newUser)

	return c.JSON(http.StatusCreated, response)
}

func timeToString(t time.Time) string {
	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}

func newCustomerResponse(c *Customer) *CustomerResponse {
	return &CustomerResponse{
		Name:         c.Name,
		Tel:          c.Tel,
		Address:      c.Address,
		ID:           c.ID,
		RegisterDate: timeToString(c.RegisterDate),
		Message:      "success",
	}
}

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/customers", CreateNewCustomer)
	e.Logger.Fatal(e.Start(":3000"))
}
