package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go-crud-echo-mysql/model"

	"fmt"
)

func GetMember(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	res := model.Member{}
	res.FirstById(id)

	return c.JSON(http.StatusOK, res)
}

func GetAllMembers(c echo.Context) error {
	m := []model.Member{}
	model.DB.Find(&m)

	return c.JSON(http.StatusOK, m)
}

func CreateMember(c echo.Context) (err error) {
	p, _ := strconv.Atoi(c.FormValue("age"))
	age := uint(p)

	member := model.Member {
		Name: c.FormValue("name"),
		Age: age,
		Bloodtype: c.FormValue("bloodtype"),
		Origin: c.FormValue("origin"),
	}

	fmt.Print("add member: ",member)
	member.Create()
	
	return c.JSON(http.StatusOK, member)
}
