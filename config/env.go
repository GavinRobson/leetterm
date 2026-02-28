package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SupabaseURL string
	SupabaseKey string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	SupabaseURL = os.Getenv("SUPABASE_URL")
	SupabaseKey = os.Getenv("SUPABASE_KEY")
}
