package controllers

import (
	"fmt"
	"goflow/model"
	"os"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AssetErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func CheckAssetPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AssetMigration() {
	godotenv.Load()
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL_STRING")), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&model.Asset{})
}

// Create Asset
func CreateAsset(c *fiber.Ctx) error {
	asset := new(model.Asset)
	asset.ID = uuid.New()

	if err := c.BodyParser(asset); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var checkexists model.CheckAssetExist
	DB.Table("assets").Select("registration_number").Where("registration_number = ?", asset.RegistrationNumber).Scan(&checkexists)
	if checkexists.RegistrationNumber == asset.RegistrationNumber {
		return c.Status(500).SendString("Asset Already Exists")
	}

	errors := ValidateAssetStruct(*asset)
	if errors != nil {
		return c.JSON(errors)
	}

	DB.Create(&asset)
	return c.JSON(&asset)
}

// Get All Assets
func GetAssets(c *fiber.Ctx) error {
	var assets []model.Asset
	DB.Find(&assets)
	return c.JSON(&assets)
}

// Count Assets
func GetAssetsCount(c *fiber.Ctx) error {
	var assetCount int64
	DB.Table("assets").Select("*").Count(&assetCount)
	// return c.JSON(&assetCount)
	return c.JSON(fiber.Map{"status": 200, "coun": &assetCount})
}

// Get Asset by ID
func GetAsset(c *fiber.Ctx) error {
	id := c.Params("id")
	var asset model.Asset
	DB.Find(&asset, id)
	return c.JSON(&asset)
}

// Validate asset before Posting
func ValidateAssetStruct(asset model.Asset) []*AssetErrorResponse {
	var errors []*AssetErrorResponse
	validate := validator.New()
	err := validate.Struct(asset)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element AssetErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// Delete Asset by ID
func DeleteAsset(c *fiber.Ctx) error {
	id := c.Params("id")
	var asset model.Asset
	DB.Table("assets").Where("id =?", id).Scan(&asset)
	if asset.RegistrationNumber == "" {
		return c.Status(500).SendString("Asset not available")
	}
	DB.Table("assets").Where("id =?", id).Unscoped().Delete(&asset)
	return c.SendString("Asset has been deleted")
}

// PUT Asset by ID
func UpdateAsset(c *fiber.Ctx) error {
	id := c.Params("id")
	var asset model.Asset
	DB.Table("assets").Where("id =?", id).Scan(&asset)
	if asset.RegistrationNumber == "" {
		return c.Status(500).SendString("Asset not available")
	}
	if err := c.BodyParser(&asset); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&asset)
	return c.JSON(&asset)
}
