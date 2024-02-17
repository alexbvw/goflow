package controllers

import (
	"fmt"
	"goflow/model"
	"goflow/util"
	"os"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IdentityErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Login Identity
func Login(c *fiber.Ctx) error {
	identity := new(model.Identity)
	if err := c.BodyParser(identity); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "identity not found"})
	}
	type Token struct {
		FullName string
		Token    string
	}
	var checkexists model.CheckIdentityExist
	var identity_token Token
	DB.Table("identities").Select("*").Where("phone_number = ?", identity.PhoneNumber).Scan(&checkexists)
	match := CheckPasswordHash(identity.PinCode, checkexists.PinCode)
	if checkexists.Role == "USER" && checkexists.PhoneNumber == identity.PhoneNumber && match {
		t, err := util.GenerateNewUserAccessToken(model.CheckIdentityExist(checkexists))

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		identity_token.Token = t
		identity_token.FullName = checkexists.FullName

		return c.JSON(&identity_token)
	} else if checkexists.Role == "ADMIN" && checkexists.PhoneNumber == identity.PhoneNumber && match {
		t, err := util.GenerateNewAdminAccessToken(model.CheckIdentityExist(checkexists))

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		identity_token.Token = t
		identity_token.FullName = checkexists.FullName

		return c.JSON(&identity_token)
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": fiber.StatusUnauthorized, "message": "identity not found"})
}

// Register Identity
func Register(c *fiber.Ctx) error {
	identity := new(model.Identity)
	identity.ID = uuid.New()

	if err := c.BodyParser(identity); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var checkexists model.CheckIdentityExist
	DB.Table("identities").Select("phone_number").Where("phone_number = ?", identity.PhoneNumber).Scan(&checkexists)
	if checkexists.PhoneNumber == identity.PhoneNumber {
		return c.Status(500).JSON(fiber.Map{"status": 500, "message": "identity already exists"})
	}
	fmt.Println(identity.PinCode)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(identity.PinCode), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(500).SendString("Could not HASH Password")
	}

	errors := ValidateIdentityStruct(*identity)
	if errors != nil {
		return c.JSON(errors)
	}

	identity.PinCode = string(hashedPassword)
	DB.Create(&identity)
	identity.PinCode = ""
	return c.Status(200).JSON(fiber.Map{"status": 200, "message": "identity registered", "identity": &identity})
}

func IdentityMigration() {
	godotenv.Load()
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL_STRING")), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&model.Identity{})
}

// Create Identity
func CreateIdentity(c *fiber.Ctx) error {
	identity := new(model.Identity)
	identity.ID = uuid.New()

	if err := c.BodyParser(identity); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var checkexists model.CheckIdentityExist
	DB.Table("identities").Select("phone_number").Where("phone_number = ?", identity.PhoneNumber).Scan(&checkexists)

	if checkexists.PhoneNumber == identity.PhoneNumber {
		return c.Status(500).JSON(fiber.Map{"status": 500, "message": "identity already exists"})
	}

	errors := ValidateIdentityStruct(*identity)
	if errors != nil {
		return c.JSON(errors)
	}

	DB.Create(&identity)
	return c.Status(200).JSON(fiber.Map{"status": 200, "message": "identity successfully created", "identity": identity})
}

// Get All Identities
func GetIdentities(c *fiber.Ctx) error {
	var identities []model.Identity
	DB.Find(&identities)
	return c.JSON(&identities)
}

// Count Identities
func GetIdentitiesCount(c *fiber.Ctx) error {
	var identityCount int64
	DB.Table("identities").Select("*").Count(&identityCount)
	// return c.JSON(&identityCount)
	return c.JSON(fiber.Map{"status": 200, "identitycount": &identityCount})
}

// Get Identity by ID
func GetIdentity(c *fiber.Ctx) error {
	id := c.Params("id")
	var identity model.Identity
	DB.Table("identities").Where("id =?", id).Scan(&identity)
	return c.JSON(&identity)
}

// Validate identity before Posting
func ValidateIdentityStruct(identity model.Identity) []*IdentityErrorResponse {
	var errors []*IdentityErrorResponse
	validate := validator.New()
	err := validate.Struct(identity)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element IdentityErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// Delete Identity by ID
func DeleteIdentity(c *fiber.Ctx) error {
	id := c.Params("id")
	var identity model.Identity
	DB.Table("identities").Where("id =?", id).Scan(&identity)
	if identity.PhoneNumber == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "identity not found",
		})
	}
	DB.Table("identities").Where("id =?", id).Unscoped().Delete(&identity)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "identity has been deleted",
	})
}

// PUT Identity by ID
func UpdateIdentity(c *fiber.Ctx) error {
	id := c.Params("id")
	var identity model.Identity
	DB.Table("identities").Where("id =?", id).Scan(&identity)
	if identity.PhoneNumber == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "identity not found",
		})
	}
	if err := c.BodyParser(&identity); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&identity)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"identity": &identity,
		"message":  "identity has been updated",
	})
}
