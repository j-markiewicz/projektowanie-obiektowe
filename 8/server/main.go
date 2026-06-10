package main

import (
	"crypto/rand"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("/products", listProducts)
	e.GET("/products/:id", readProduct)

	e.POST("/login", login)
	e.POST("/signup", signup)

	if err := e.Start(":8000"); err != nil {
		panic("failed to start server:\n" + err.Error())
	}
}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
}

var products = map[string]Product{
	"1": {"Truskawki", "opakowanie 500g", 699},
	"2": {"Ogórki", "1kg, świeże", 899},
	"3": {"Marchewki", "1kg luzem", 399},
}

type User struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Method string `json:"method"`
	Token  string `json:"token"`
}

var users = map[string]User{}

var tokens = map[string]string{}

// GET /products
func listProducts(c *echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

// GET /products/:id
func readProduct(c *echo.Context) error {
	return c.JSON(http.StatusOK, products[c.Param("id")])
}

// POST /login
func login(c *echo.Context) error {
	email := c.FormValue("email")
	method := c.FormValueOr("provider", "password")

	if users[email].Email != email {
		return c.String(http.StatusForbidden, "invalid email address")
	}

	if users[email].Method != method {
		return c.String(http.StatusForbidden, "invalid authentication provider")
	}

	if method == "password" {
		password := c.FormValue("password")

		if users[email].Token != password {
			return c.String(http.StatusForbidden, "invalid password")
		}
	} else {
		return c.NoContent(http.StatusForbidden)
	}

	token := rand.Text()
	tokens[token] = email
	c.SetCookie(&http.Cookie{Name: "auth"})
	return c.String(http.StatusOK, users[email].Name)
}

// POST /signup
func signup(c *echo.Context) error {
	email := c.FormValue("email")
	name := c.FormValue("name")
	password := c.FormValue("password")

	_, err := mail.ParseAddress(email)
	if err != nil {
		return c.String(http.StatusForbidden, "invalid email address")
	}

	if users[email].Email != "" {
		return c.String(http.StatusForbidden, "user already exists")
	}

	users[email] = User{
		Email:  email,
		Name:   name,
		Method: "password",
		Token:  password,
	}

	return c.NoContent(http.StatusOK)
}
