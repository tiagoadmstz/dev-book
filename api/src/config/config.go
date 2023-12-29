package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// MongoDB string connection
	StringDatabaseConnection = ""
	// MongoDB database name
	DatabaseName = ""
	// MongoDB cluster name
	ClusterName = ""
	// API running port
	Port = 0
)

// Charge initialize environment variables
func Charge() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err = godotenv.Load(); err != nil {
		Port = 9000
	}

	StringDatabaseConnection = fmt.Sprintf("mongodb://%s:%s@%s:27017/?retryWrites=true&w=majority",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_CLUSTER"),
	)

	DatabaseName = os.Getenv("DB_NAME")
}
