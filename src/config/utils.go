package config

import (
	"log"
	"os"
	"strconv"
)

var (
	SecretKey      []byte
	Address        string
	MongodbURI     string
	IsClustered    bool
	TestingEnabled bool
)

// This function initializes and loads configuration variables from environment variables in a Go
// application.
func initialize() {

	secretKeyEnv := os.Getenv("JWT_SECRET")
	if secretKeyEnv != "" {
		SecretKey = []byte(secretKeyEnv)
	} else {
		SecretKey = []byte("test_local")
	}

	Address = os.Getenv("ADDRESS")
	if Address == "" {
		Address = "127.0.0.1:6969"
	}

	isClustered := os.Getenv("IS_CLUSTERED")
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

	testingEnabled := os.Getenv("TESTING_ENABLED")
	if testingEnabled != "" {
		var err error
		TestingEnabled, err = strconv.ParseBool(testingEnabled)
		if err != nil {
			TestingEnabled = false
			log.Println("Config: Failed to parse TESTING_ENABLED as boolean, using default value false")
		}
	} else {
		TestingEnabled = false
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
	if TestingEnabled {
		log.Printf("Config:    SecretKey: %s\n", SecretKey)
	}
	log.Printf("Config:    Address: %s\n", Address)
	log.Printf("Config:    IsClustered: %v\n", IsClustered)
	log.Printf("Config:    TestingEnabled: %v\n", TestingEnabled)
	log.Printf("Config:    MongodbURI: %s\n", MongodbURI)
}

func init() {
	initialize()
}

func Reset() {
	initialize()
}
