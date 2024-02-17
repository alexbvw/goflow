package controllers

import (
	"fmt"
	"goflow/model"
	"os"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ChatroomErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ChatroomMigration() {
	godotenv.Load()
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL_STRING")), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&model.Chatroom{})
}

// Get All Chatrooms
func GetChatrooms(c *fiber.Ctx) error {
	var chatrooms []model.Chatroom
	DB.Find(&chatrooms)
	return c.JSON(&chatrooms)
}

// Count Chatrooms
func GetChatroomsCount(c *fiber.Ctx) error {
	var chatrooms []model.Chatroom
	var chatroomCount int64
	DB.Find(&chatrooms).Count(&chatroomCount)
	// return c.JSON(&chatroomCount)
	return c.JSON(fiber.Map{"status": 200, "chatroomcount": &chatroomCount})
}

// Get Chatroom by ID
func GetChatroom(c *fiber.Ctx) error {
	id := c.Params("id")
	var chatroom model.Chatroom
	DB.Raw(`SELECT * FROM chatrooms WHERE id = $1`, id).Scan(&chatroom)
	return c.JSON(&chatroom)
}

// Get Chatroom by UserID
func GetChatroomsbyIdentityId(c *fiber.Ctx) error {
	id := c.Params("id")
	var chatrooms []model.Chatroom
	DB.Raw(`SELECT * FROM Chatrooms WHERE $1 = ANY(identities)`, id).Scan(&chatrooms)
	return c.JSON(&chatrooms)
}

// Create Chatroom
func CreateChatroom(c *fiber.Ctx) error {
	chatroom := new(model.Chatroom)
	chatroom.ID = uuid.New()
	if err := c.BodyParser(chatroom); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	errors := ValidateChatroomStruct(*chatroom)
	if errors != nil {
		return c.JSON(errors)
	}
	DB.Create(&chatroom)
	return c.JSON(&chatroom)
}

// Delete Chatroom by ID
func DeleteChatroom(c *fiber.Ctx) error {
	id := c.Params("id")
	var chatroom model.Chatroom
	DB.Raw(`SELECT * FROM chatrooms WHERE id = $1`, id).Scan(&chatroom)
	if chatroom.ChatroomName == "" {
		return c.Status(500).SendString("Chatroom not available")
	}
	DB.Delete(&chatroom).Where("id =? ", chatroom.ID)
	return c.SendString("Chatroom has been deleted")
}

// PATCH Chatroom by ID
func UpdateChatroom(c *fiber.Ctx) error {
	id := c.Params("id")
	var chatroom model.Chatroom
	DB.First(&chatroom, id)
	if chatroom.ChatroomName == "" {
		return c.Status(500).SendString("Chatroom not available")
	}
	if err := c.BodyParser(&chatroom); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&chatroom)
	return c.JSON(&chatroom)
}

// Validate chatroom before Posting
func ValidateChatroomStruct(chatroom model.Chatroom) []*ChatroomErrorResponse {
	var errors []*ChatroomErrorResponse
	validate := validator.New()
	err := validate.Struct(chatroom)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ChatroomErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
