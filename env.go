package best_practices_golang

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func Env() (string, string, string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("❌ Failed to get working directory: %v", err)
	}

	envPath := filepath.Join(wd, ".env")
	log.Println("🔍 Loading .env from:", envPath)

	_ = godotenv.Load(envPath)

	redis := os.Getenv("REDIS_ADDRESS")
	if redis == "" {
		log.Fatal("❌ Environment variable REDIS_ADDRESS is not set")
	}

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("❌ Environment variable MYSQL_DSN is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("❌ Environment variable PORT is not set")
	}

	return redis, dsn, port
}
