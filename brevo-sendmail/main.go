package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	BrevoAPIURL = "https://api.brevo.com/v3/smtp/email"
)

type BrevoEmailRequest struct {
	Sender      Sender      `json:"sender"`
	To          []Recipient `json:"to"`
	Subject     string      `json:"subject"`
	HTMLContent string      `json:"htmlContent"`
}

type Sender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Recipient struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GenerateActivationCode() string {
	rand.Seed(time.Now().UnixNano())

	code := rand.Intn(900000) + 100000
	return strconv.Itoa(code)
}

func LoadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

func SendActivationEmail(recipientEmail, recipientName, activationCode string) error {
	apiKey := os.Getenv("BREVO_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("API KEY is not set")
	}

	senderName := os.Getenv("SENDER_NAME")
	senderEmail := os.Getenv("SENDER_EMAIL")
	if senderName == "" || senderEmail == "" {
		return fmt.Errorf("sender information not set in environment variables")
	}

	emailReq := BrevoEmailRequest{
		Sender: Sender{
			Name:  senderName,
			Email: senderEmail,
		},
		To: []Recipient{{
			Email: recipientEmail,
			Name:  recipientName,
		},
		},
		Subject: "Activation Code",
		HTMLContent: fmt.Sprintf(`
			<html>
				<body>
					<h1>Welcome to Your Company!</h1>
					<p>Hello %s,</p>
					<p>Thank you for signing up. Please use the following code to activate your account:</p>
					<h2 style="background-color: #f0f0f0; padding: 10px; text-align: center; font-size: 24px;">%s</h2>
					<p>This code will expire in 15 minutes.</p>
					<p>If you didn't request this code, please ignore this email.</p>
					<p>Best regards,<br>Your Company Team</p>
				</body>
			</html>
		`, recipientName, activationCode),
	}

	jsonData, err := json.Marshal(emailReq)
	if err != nil {
		return fmt.Errorf("failed to marshal request to JSON: %v", err)
	}

	req, err := http.NewRequest("POST", BrevoAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	log.Printf("Response from Brevo API: %s", string(body))

	return nil
}

func main() {
	err := LoadEnvVariables()
	if err != nil {
		log.Fatalf("Failed load env variables: %v", err)
	}

	testEmail := os.Getenv("TEST_RECIPIENT_EMAIL")
	testName := os.Getenv("TEST_RECIPIENT_NAME")

	activationCode := GenerateActivationCode()

	err = SendActivationEmail(testEmail, testName, activationCode)
	if err != nil {
		log.Fatalf("Error sending activation email: %v", err)
	}

	log.Printf("Activation email sent successfully to %s with code %s", testEmail, activationCode)
}
