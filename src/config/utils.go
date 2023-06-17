package config

import (
	"log"
	"os"
	"strconv"
)

var (
	SecretKey   []byte
	Address     string
	isClustered string
	MongodbURI  string
	IsClustered bool
)

// This function initializes and loads configuration variables from environment variables in a Go
// application.
func init() {

	secretKeyEnv := os.Getenv("JWT_SECRET")
	if secretKeyEnv != "" {
		SecretKey = []byte(secretKeyEnv)
	} else {
		SecretKey = []byte("test_local")
	}

	Address = os.Getenv("ADDRESS")
	if Address == "" {
		Address = ":6969"
	}

	isClustered = os.Getenv("IS_CLUSTERED")
	if isClustered != "" {
		var err error
		IsClustered, err = strconv.ParseBool(isClustered)
		if err != nil {
			IsClustered = false
			log.Println("Config: Failed to parse IS_CLUSTERED as boolean, using default value false")
		}
	} else {
		IsClustered = false
	}

	MongoUser := os.Getenv("MONGO_USERNAME")
	MongoPassword := os.Getenv("MONGO_PASSWORD")
	MongodbURI = os.Getenv("MONGO_SERVICE")
	if MongodbURI == "" {
		log.Println("Config: MONGODB_URI environment variable not set, using local storage")
	} else {
		//MongodbURI = strings.Replace(MongodbURI, "tcp", "mongodb", 1)

		MongodbURI = "mongodb://" + MongoUser + ":" + MongoPassword + "@" + MongodbURI
	}

	// Logging the loaded configuration
	log.Println("Config: Loaded configuration:")
	log.Printf("Config:    SecretKey: %s\n", SecretKey)
	log.Printf("Config:    Address: %s\n", Address)
	log.Printf("Config:    IsClustered: %v\n", IsClustered)
	log.Printf("Config:    MongodbURI: %s\n", MongodbURI)
}
