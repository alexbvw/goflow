package controllers

import (
	"encoding/json"
	"goflow/model"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func RequestTokenHandler(c *fiber.Ctx) error {
	code := c.Query("code") // Assuming 'code' is passed as a query parameter
	var token model.Token
	client := &http.Client{}
	godotenv.Load()

	// Setup the request
	req, err := http.NewRequest("POST", os.Getenv("WEBFLOW_BASE_URL")+"/oauth/access_token?client_id="+os.Getenv("CLIENT_ID")+"&client_secret="+os.Getenv("CLIENT_SECRET")+"&code="+code+"&grant_type=authorization_code&redirect_uri="+os.Getenv("REDIRECT_URI"), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	req.Header.Add("accept-version", "1.0.0")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to read response body", "error": err.Error()})
	}

	if err := json.Unmarshal(body, &token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unmarshal response"})
	}

	// Return the response body as part of the JSON response
	return c.JSON(fiber.Map{"response": token})
}

func FetchSitesHandler(c *fiber.Ctx) error {
	client := &http.Client{}

	// Retrieve the token from the request headers
	token := c.Get("Authorization")

	// If the token is empty, return an error response
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Authorization token is required"})
	}

	req, err := http.NewRequest("GET", os.Getenv("WEBFLOW_BASE_URL")+"/v2/sites", nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Use the token from the request headers for authorization
	req.Header.Add("Authorization", token)
	req.Header.Add("accept-version", "1.0.0")

	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	var sitesResponse model.SitesResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to read response body", "error": err.Error()})
	}

	if err := json.Unmarshal(body, &sitesResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unmarshal response"})
	}

	return c.JSON(sitesResponse)
}

func FetchCollectionsHandler(c *fiber.Ctx) error {
	siteId := c.Query("siteId")
	if siteId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Site ID is required"})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Authorization token is required"})
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("WEBFLOW_BASE_URL")+"/v2/sites/"+siteId+"/collections", nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create request"})
	}

	req.Header.Add("Authorization", token)
	req.Header.Add("accept-version", "1.0.0")

	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to make request"})
	}
	defer resp.Body.Close()

	var collectionsResponse model.CollectionsResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read response body"})
	}

	if err := json.Unmarshal(body, &collectionsResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unmarshal response"})
	}

	return c.JSON(collectionsResponse)
}

func FetchCollectionItemsHandler(c *fiber.Ctx) error {
	collectionId := c.Query("collectionId")
	if collectionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Collection ID is required"})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Authorization token is required"})
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("WEBFLOW_BASE_URL")+"/v2/collections/"+collectionId+"/items", nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create request"})
	}

	req.Header.Add("Authorization", token)
	req.Header.Add("accept-version", "1.0.0")

	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to make request"})
	}
	defer resp.Body.Close()

	var itemsResponse model.CollectionItemsResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read response body"})
	}

	if err := json.Unmarshal(body, &itemsResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unmarshal response"})
	}

	return c.JSON(itemsResponse)
}
