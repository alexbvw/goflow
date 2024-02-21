package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goflow/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MessageErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Client struct {
	chatroomId string
	events     chan *MessageChannel
}
type MessageChannel struct {
	// User  uint
	Messages []model.Message
}

func MessageMigration() {
	godotenv.Load()
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL_STRING")), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&model.Message{})
}

// SSE Messages
func Handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}

func ChatroomMessagesHandler(w http.ResponseWriter, r *http.Request) {

	chatroomId, err := r.URL.Query()["chatroomId"]

	if !err || len(chatroomId[0]) < 1 {
		log.Println("URL Param 'chatroomId' is missing")
		return
	}

	client := &Client{chatroomId: chatroomId[0], events: make(chan *MessageChannel, 10)}

	go updatechatroomMessages(client)
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	timeout := time.After(2 * time.Second)

	select {
	case ev := <-client.events:
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(ev)
		fmt.Fprintf(w, "data: %v\n\n", buf.String())
	case <-timeout:
		fmt.Fprintf(w, ": nothing to sent\n\n")
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func updatechatroomMessages(client *Client) {
	var messages []model.Message
	DB.Raw(`SELECT * FROM Messages WHERE chatroom_id = $1`, client.chatroomId).Scan(&messages)

	if err != nil {
		println(err)
	}

	for {
		db := &MessageChannel{
			Messages: messages,
		}

		client.events <- db
	}
}

// Get All Messages
func GetMessages(c *fiber.Ctx) error {
	var messages []model.Message
	DB.Find(&messages)
	return c.JSON(&messages)
}

// Count Messages
func GetMessagesCount(c *fiber.Ctx) error {
	var messages []model.Message
	var messageCount int64
	DB.Find(&messages).Count(&messageCount)
	// return c.JSON(&messageCount)
	return c.JSON(fiber.Map{"status": 200, "messagecount": &messageCount})
}

// Get Message by ID
func GetMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	var message model.Message
	DB.Find(&message, id)
	return c.JSON(&message)
}

// Create Message
func CreateMessage(c *fiber.Ctx) error {
	message := new(model.Message)
	message.ID = uuid.New()
	if err := c.BodyParser(message); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	errors := ValidateMessageStruct(*message)
	if errors != nil {
		return c.JSON(errors)
	}
	DB.Create(&message)
	return c.JSON(&message)
}

// Delete Message by ID
func DeleteMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	var message model.Message
	DB.First(&message, id)
	if message.Message == "" {
		return c.Status(500).SendString("Message not available")
	}
	DB.Delete(&message)
	DB.Unscoped().Delete(&message, id)
	return c.SendString("Message has been deleted")
}

// PATCH Message by ID
func UpdateMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	var message model.Message
	DB.First(&message, id)
	if message.Message == "" {
		return c.Status(500).SendString("Message not available")
	}
	if err := c.BodyParser(&message); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&message)
	return c.JSON(&message)
}

// Validate message before Posting
func ValidateMessageStruct(message model.Message) []*MessageErrorResponse {
	var errors []*MessageErrorResponse
	validate := validator.New()
	err := validate.Struct(message)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element MessageErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
