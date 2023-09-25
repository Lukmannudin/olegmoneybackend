package config

import (
	"log"
	"os"

	supa "github.com/nedpals/supabase-go"

	"github.com/joho/godotenv"
)

func GetConnectionSupabse() *supa.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	return supabase
}
