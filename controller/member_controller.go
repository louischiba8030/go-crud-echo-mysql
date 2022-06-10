package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/google/uuid"

	"go-crud-echo-mysql/model"

	"fmt"
)

func GetMember(c echo.Context) error {
	m := model.Member{}

	if c_uuid, err := uuid.Parse(c.Param("id")); err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println(c_uuid)
		m.FirstById(c_uuid)
	}
	fmt.Printf("m.ID = %v", m.ID)
	fmt.Printf("m.Name = %v", m.Name)

	return c.JSON(http.StatusOK, m)
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

func UpdateMember (c echo.Context) (err error) {
	p, _ := strconv.Atoi(c.FormValue("age"))
	age := uint(p)

	c_uuid, _ := uuid.Parse(c.Param("id"))
	fmt.Printf("c_uuid = ", c_uuid)
	m := model.Member {
		Name: c.FormValue("name"),
		Age: age,
		Bloodtype: c.FormValue("bloodtype"),
		Origin: c.FormValue("origin"),
	}
	fmt.Println(c.FormValue("age"))
	m.Updates(c_uuid)

	return c.JSON(http.StatusOK, m)
}

func DeleteMember (c echo.Context) (err error) {
	m := model.Member{}

	if c_uuid, err := uuid.Parse(c.Param("id")); err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println(c_uuid)
		m.DeleteById(c_uuid)
	}

	return c.JSON(http.StatusOK, m)
}
