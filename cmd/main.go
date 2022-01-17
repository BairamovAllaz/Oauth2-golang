package main

import (
	"Project/Server"
	"Project/pkg/handler"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	// "fmt"
)

func main() { 

	GoogleAuth := &oauth2.Config{
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint: google.Endpoint,
		Scopes: []string{"https://googleapis.com/auth/userinfo.email","https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL: "http://localhost:8000/callback",
	}

	//dotenv confugration
	
	if err := godotenv.Load(filepath.Join(".",".env"));err != nil { 
		log.Fatalf("Error %s", err.Error());
	}


	handler := handler.NewHandler(GoogleAuth)
	srv := new(server.Server)
	if err := srv.Run("8000",handler.InitRoutes());err != nil { 
		log.Fatalf("Error %s", err.Error())
	}
}