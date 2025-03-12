package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/milan-kovac/database"
	"github.com/milan-kovac/lead"
)

func setRoutes(app *fiber.App){
app.Get("/api/v1/lead",lead.GetLeads)
app.Get("/api/v1/lead/:id", lead.GetLead)
app.Post("/api/v1/lead", lead.NewLeead)
app.Delete("/api/v1/lead/:id" ,lead.DeleteLead)
}

func initDatabase(){
  var err error 
  database.DB_CONNECTION, err = gorm.Open("sqllite3", "leads.db")

  if err != nil {
	panic("Failed to connect to database.")
  }

  fmt.Println("Connection opened to database.")

  database.DB_CONNECTION.AutoMigrate(&lead.Lead{})

  fmt.Println("Database migrated.")
}

func main(){
	app := fiber.New()
	initDatabase()
	setRoutes(app)


	app.Listen("3000")
	defer database.DB_CONNECTION.Close()
}