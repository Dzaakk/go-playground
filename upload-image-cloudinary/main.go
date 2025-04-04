package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		log.Fatal("Missing required environment variables. Please check your .env file.")
	}

	cloud, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}

	imagePath := os.Getenv("IMAGE_PATH")
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	uploadResult, err := cloud.Upload.Upload(
		context.Background(),
		file,
		uploader.UploadParams{
			Folder: os.Getenv("CLOUDINARY_FOLDER"),
		},
	)

	if err != nil {
		log.Fatalf("Failed to upload image: %v", err)
	}

	fmt.Println("Upload successful!")
	fmt.Println("URL:", uploadResult.SecureURL)
	fmt.Println("Public ID:", uploadResult.PublicID)
}
