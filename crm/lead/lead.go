package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/milan-kovac/database"
)

type Lead struct {
	gorm.Model
	Name string `json:"name"`
	Company string `json:"company"`
	Email string  `json:"email"`
	Phone int `json:"phone"`
}

func GetLeads(context *fiber.Ctx) error  {
	db:=database.DB_CONNECTION
	var leads []Lead
	
	db.Find(&leads)

	return context.JSON(leads)
}

func GetLead(context *fiber.Ctx) error {
	id := context.Params("id")
	db:=database.DB_CONNECTION

	var lead Lead
	db.Find(&lead, id)

	return context.JSON(lead)
}

func NewLead(context *fiber.Ctx) error {
	db:=database.DB_CONNECTION

	lead := new(Lead)

	if err:= context.BodyParser(lead); err !=nil{
		return context.Status(503).Send([]byte(err.Error()))
	}

	db.Create(lead)

	return context.JSON(lead)
}

func DeleteLead(context *fiber.Ctx) error {
	db:=database.DB_CONNECTION
	id := context.Params("id")

	var lead Lead
	db.First(&lead, id)

	if  lead.Name == "" {
		return context.Status(400).SendString("Lead not found.")
	}
     
	db.Delete(&lead)

	return  context.Status(200).SendString("Lead deleted.")
}