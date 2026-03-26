package main

import (
	"backend/config"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("warning: .env file not found")
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.GET("/", func(c *echo.Context) error {
		return c.String(200, "Backend connected to PostgreSQL")
	})

	e.GET("/products", func(c *echo.Context) error {
		products, err := config.GetProducts(db)
		if err != nil {
			log.Println("get products error:", err)

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(), 
			})
		}
		return c.JSON(200, products)
	})

	e.GET("/categories", func(c *echo.Context) error {
		categories, err := config.GetCategories(db)
		if err != nil {
			log.Println("get categories error:", err)

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(), 
			})
		}
		return c.JSON(200, categories)
	})

	log.Println("Server running at :8080")
	log.Fatal(e.Start(":8080"))
}
