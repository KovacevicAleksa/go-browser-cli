package AI

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const apiURL = "https://openrouter.ai/api/v1/chat/completions"

type RequestBody struct {
	Model       string              `json:"model"`
	Messages    []map[string]string `json:"messages"`
	MaxTokens   int                 `json:"max_tokens"`
	Temperature float64             `json:"temperature"`
}

type ResponseBody struct {
	Choices []struct {
		Message map[string]string `json:"message"`
	} `json:"choices"`
}

func ChatGPT(input string) string {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR: Error loading .env file")
		return ""
	}

	// Get API key from environment variable
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		log.Println("ERROR: API key is missing")
		return ""
	}

	// Use the appropriate model
	model := "google/gemini-2.0-flash-exp:free"

	// Prepare the messages array
	messages := []map[string]string{
		{"role": "user", "content": input},
	}

	// Create the request body
	requestBody := RequestBody{
		Model:       model,
		Messages:    messages,
		MaxTokens:   100,
		Temperature: 0.7,
	}

	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Println("ERROR: Error marshaling request:", err)
		return ""
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("ERROR: Error creating request:", err)
		return ""
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ERROR: Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	// Parse the response
	var responseBody ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Println("ERROR: Error decoding response:", err)
		return ""
	}

	// Return the first choice content
	if len(responseBody.Choices) > 0 {
		return responseBody.Choices[0].Message["content"]
	}
	return "No response from API"
}
