//package router
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-crud-echo-mysql/controller"

//	"fmt"
)

/*func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Golang World!")
	})

	e.GET("/api/member/:id", controller.GetMember)

	return e
}
*/
func main() {
	// Create instance
	e := echo.New()

	// Set middleware
	e.Use(middleware.CORS())
	e.Use(middleware.BodyDump(bodyDumpHandler))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Golang World!")
	})
	e.GET("/api/members/", controller.GetAllMembers)
	e.POST("/api/members/create", controller.CreateMember)
	e.GET("/api/members/:id", controller.GetMember)
	e.PUT("/api/members/:id", controller.UpdateMember)
	e.DELETE("/api/members/:id", controller.DeleteMember)

	// Launch server at port 8003
	e.Logger.Fatal(e.Start(":8003"))
}

// Define handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Golang World!")
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
//	fmt.Printf("Request body: %v\n", string(reqBody))
//	fmt.Printf("Response body: %v\n", string(resBody))
}