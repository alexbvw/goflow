package controllers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"goflow/model"
	"io"
	"mime/multipart"
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

func GetAuthorizedUserInfo(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	// If the token is empty, return an error response
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Authorization token is required"})
	}

	client := &http.Client{}

	// Setup the request
	req, err := http.NewRequest("GET", os.Getenv("WEBFLOW_BASE_URL")+"/v2/token/authorized_by", nil)
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

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to read response body", "error": err.Error()})
	}

	// Return the response body as part of the JSON response
	return c.JSON(fiber.Map{"response": string(body)})
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

func FetchSiteHandler(c *fiber.Ctx) error {
	siteId := c.Params("id")
	if siteId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Site ID is required"})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Authorization token is required"})
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("WEBFLOW_BASE_URL")+"/v2/sites/"+siteId, nil)
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

	var siteResponse model.SiteResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to read response body", "error": err.Error()})
	}

	if err := json.Unmarshal(body, &siteResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unmarshal response"})
	}

	return c.JSON(siteResponse)

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

func FetchCollectionHandler(c *fiber.Ctx) error {
	collectionId := c.Params("id")
	if collectionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Collection ID is required"})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Authorization token is required"})
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("WEBFLOW_BASE_URL")+"/v2/collections/"+collectionId, nil)
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

	if resp.StatusCode != http.StatusOK {
		return c.Status(resp.StatusCode).JSON(fiber.Map{"error": "Request failed with status " + resp.Status})
	}

	var collectionResponse model.Collection
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to read response body", "error": err.Error()})
	}

	// Debugging: Print raw response
	fmt.Println("Raw Response Body:", string(body))

	if err := json.Unmarshal(body, &collectionResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unmarshal response", "details": err.Error()})
	}

	// Debugging: Print unmarshalled response
	fmt.Println("Unmarshalled Response:", collectionResponse)

	return c.JSON(collectionResponse)
}

func FetchCollectionItemHandler(c *fiber.Ctx) error {
	collectionId := c.Params("collectionId")
	itemId := c.Params("itemId")
	if collectionId == "" || itemId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Collection ID or Item ID is required"})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Authorization token is required"})
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("WEBFLOW_BASE_URL")+"/v2/collections/"+collectionId+"/items/"+itemId, nil)
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

	if resp.StatusCode != http.StatusOK {
		return c.Status(resp.StatusCode).JSON(fiber.Map{"error": "Request failed with status " + resp.Status})
	}

	var itemResponse model.CollectionItem
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to read response body", "error": err.Error()})
	}

	// Debugging: Print raw response
	fmt.Println("Raw Response Body:", string(body))

	if err := json.Unmarshal(body, &itemResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unmarshal response", "details": err.Error()})
	}

	// Debugging: Print unmarshalled response
	fmt.Println("Unmarshalled Response:", itemResponse)

	return c.JSON(itemResponse)
}

func UpdateCollectionItemsHandler(c *fiber.Ctx) error {
	// 1. Grab path parameters
	collectionId := c.Params("collectionId")
	if collectionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Collection ID is required",
		})
	}
	fmt.Println("Collection ID:", collectionId)
	// itemId := c.Params("itemId")
	// if itemId == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Item ID is required",
	// 	})
	// }

	// 2. Authorization header
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Authorization token is required",
		})
	}

	// 3. Read the raw request body
	rawBody := c.Body() // <-- This is the raw JSON from the client

	// 4. Build the Webflow API endpoint
	endpoint := fmt.Sprintf("%s/v2/collections/%s/items",
		os.Getenv("WEBFLOW_BASE_URL"),
		collectionId,
	)

	// 5. Create a new PATCH request
	req, err := http.NewRequest(http.MethodPatch, endpoint, bytes.NewBuffer(rawBody))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create request: %v", err),
		})
	}

	// 6. Set headers
	req.Header.Set("Authorization", token)
	req.Header.Set("accept-version", "1.0.0")
	req.Header.Set("Content-Type", "application/json")

	// 7. Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Request error: %v", err),
		})
	}
	defer resp.Body.Close()

	// 8. If Webflow API did not return 200, try to read body for error details
	if resp.StatusCode != http.StatusOK {
		var errResp map[string]interface{}
		if decodeErr := json.NewDecoder(resp.Body).Decode(&errResp); decodeErr == nil {
			return c.Status(resp.StatusCode).JSON(errResp)
		}
		return c.Status(resp.StatusCode).JSON(fiber.Map{"error": "Request failed"})
	}

	// 9. If successful, decode the updated item from Webflow
	var updatedItem map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&updatedItem); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to decode response: %v", err),
		})
	}

	// 10. Return the updated item to the client
	return c.Status(fiber.StatusOK).JSON(updatedItem)
}

func PublishCollectionItemHandler(c *fiber.Ctx) error {
	// 1. Grab path parameters
	collectionId := c.Params("collectionId")
	if collectionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Collection ID is required",
		})
	}

	// 2. Authorization header
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Authorization token is required",
		})
	}

	// 3. Read the raw request body
	rawBody := c.Body() // <-- This is the raw JSON from the client

	// 4. Build the Webflow API endpoint
	endpoint := fmt.Sprintf("%s/v2/collections/%s/items/publish",
		os.Getenv("WEBFLOW_BASE_URL"),
		collectionId,
	)

	// 5. Create a new PATCH request
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(rawBody))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create request: %v", err),
		})
	}

	// 6. Set headers
	req.Header.Set("Authorization", token)
	req.Header.Set("accept-version", "1.0.0")
	req.Header.Set("Content-Type", "application/json")

	// 7. Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Request error: %v", err),
		})
	}
	defer resp.Body.Close()

	// 8. If Webflow API did not return 200, try to read body for error details
	if resp.StatusCode != http.StatusOK {
		var errResp map[string]interface{}
		if decodeErr := json.NewDecoder(resp.Body).Decode(&errResp); decodeErr == nil {
			return c.Status(resp.StatusCode).JSON(errResp)
		}
		return c.Status(resp.StatusCode).JSON(fiber.Map{"error": "Request failed"})
	}

	// 9. If successful, decode the updated item from Webflow
	var updatedItem map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&updatedItem); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to decode response: %v", err),
		})
	}

	// 10. Return the updated item to the client
	return c.Status(fiber.StatusOK).JSON(updatedItem)
}

func UploadAssetHandler(c *fiber.Ctx) error {
	siteID := c.Params("siteId")
	if siteID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Site ID is required",
		})
	}

	// Webflow requires "Authorization: Bearer <token>"
	// If your token does not include "Bearer ", add it manually.
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Authorization token is required",
		})
	}

	// 1) Read the file from the request (multipart/form-data)
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please provide a valid file in the 'file' field",
		})
	}

	// Read the file content into memory for hashing + later re-upload
	originalFile, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not open uploaded file",
		})
	}
	defer originalFile.Close()

	fileBytes, err := io.ReadAll(originalFile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not read uploaded file",
		})
	}

	// 2) Compute MD5 hash of the file
	hasher := md5.New()
	hasher.Write(fileBytes)
	md5Hash := hex.EncodeToString(hasher.Sum(nil))

	// 3) Step One: Request S3 upload details from Webflow (top-level fileName + fileHash)
	createAssetReqBody := map[string]interface{}{
		"fileName": fileHeader.Filename, // e.g. "my-image.png"
		"fileHash": md5Hash,             // e.g. "3c7d87c9575702bc3b1e991f4d3c638e"
	}
	reqData, err := json.Marshal(createAssetReqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to marshal request body for Webflow: %v", err),
		})
	}

	endpoint := fmt.Sprintf("%s/v2/sites/%s/assets", os.Getenv("WEBFLOW_BASE_URL"), siteID)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(reqData))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create request for Webflow: %v", err),
		})
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept-Version", "1.0.0")
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Request error (Webflow create asset): %v", err),
		})
	}
	defer resp.Body.Close()

	// *** Accept any 2xx as success, else treat as error ***
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp map[string]interface{}
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error":  "Webflow create asset request failed",
			"detail": errResp,
		})
	}

	// 4) Parse the JSON from Webflow's response (top-level fields)
	var wfResponse struct {
		ID            string            `json:"id"`
		UploadUrl     string            `json:"uploadUrl"`
		UploadDetails map[string]string `json:"uploadDetails"`
		AssetUrl      string            `json:"assetUrl"`
		HostedUrl     string            `json:"hostedUrl"`
		// etc.
	}
	if err := json.NewDecoder(resp.Body).Decode(&wfResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to decode Webflow response: %v", err),
		})
	}

	// 5) Step Two: Upload the actual file bytes to the returned "uploadUrl"
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add each field from "uploadDetails"
	for key, val := range wfResponse.UploadDetails {
		_ = writer.WriteField(key, val)
	}

	// Add the file contents under form field "file"
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create form file field: %v", err),
		})
	}
	_, err = part.Write(fileBytes)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to copy file content: %v", err),
		})
	}

	if err := writer.Close(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to close multipart writer: %v", err),
		})
	}

	s3Req, err := http.NewRequest(http.MethodPost, wfResponse.UploadUrl, &buf)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create S3 request: %v", err),
		})
	}
	s3Req.Header.Set("Content-Type", writer.FormDataContentType())

	s3Resp, err := httpClient.Do(s3Req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error uploading to S3: %v", err),
		})
	}
	defer s3Resp.Body.Close()

	// Webflow's S3 returns 201 on successful file upload
	if s3Resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(s3Resp.Body)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "S3 upload failed; expected 201",
			"details": string(bodyBytes),
		})
	}

	// 6) If the file was uploaded successfully to S3, Webflow finalizes the asset automatically.
	data := fiber.Map{
		"alt":    nil,                 // or fill this in if needed
		"fileId": wfResponse.ID,       // from "id"
		"url":    wfResponse.AssetUrl, // or "hostedUrl"
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "successfully uploaded file",
		"data":    data,
	})
}
