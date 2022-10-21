package main

import (
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

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
