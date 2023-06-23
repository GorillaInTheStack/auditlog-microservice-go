package config_test

import (
	"os"
	"testing"

	"auditlog/config"

	qt "github.com/frankban/quicktest"
)

func TestConfigInitialization(t *testing.T) {

	c := qt.New(t)

	// Set environment variables for testing
	os.Setenv("JWT_SECRET", "test_secret")
	os.Setenv("ADDRESS", "localhost:8080")
	os.Setenv("IS_CLUSTERED", "true")
	os.Setenv("TESTING_ENABLED", "false")
	os.Setenv("MONGO_USERNAME", "test_user")
	os.Setenv("MONGO_PASSWORD", "test_password")
	os.Setenv("MONGO_SERVICE", "mongodb://localhost:27017")

	// Perform initialization
	config.Reset()

	// Assert the initialized values
	c.Assert("test_secret", qt.Equals, string(config.SecretKey),
		qt.Commentf("Config Test: Expected value %v, got %v", "test_secret", string(config.SecretKey)))

	c.Assert("localhost:8080", qt.Equals, config.Address,
		qt.Commentf("Config Test: Expected value %v, got %v", "localhost:8080", config.Address))

	c.Assert(config.IsClustered, qt.IsTrue,
		qt.Commentf("Config Test: Expected value %v, got %v", true, config.IsClustered))

	c.Assert(config.TestingEnabled, qt.IsFalse,
		qt.Commentf("Config Test: Expected value %v, got %v", false, config.TestingEnabled))

	c.Assert("mongodb://test_user:test_password@mongodb://localhost:27017", qt.Equals, config.MongodbURI,
		qt.Commentf("Config Test: Expected value %v, got %v", "mongodb://test_user:test_password@mongodb://localhost:27017", config.MongodbURI))

}

func TestConfigInitialization_DefaultValues(t *testing.T) {

	c := qt.New(t)

	// Set environment variables for testing
	os.Setenv("JWT_SECRET", "")
	os.Setenv("ADDRESS", "")
	os.Setenv("IS_CLUSTERED", "")
	os.Setenv("TESTING_ENABLED", "")
	os.Setenv("MONGO_USERNAME", "")
	os.Setenv("MONGO_PASSWORD", "")
	os.Setenv("MONGO_SERVICE", "")

	// Perform initialization
	config.Reset()

	// Assert the default values
	// Assert the initialized values
	c.Assert("test_local", qt.Equals, string(config.SecretKey),
		qt.Commentf("Config Test: Expected value %v, got %v", "test_local", string(config.SecretKey)))

	c.Assert("127.0.0.1:6969", qt.Equals, config.Address,
		qt.Commentf("Config Test: Expected value %v, got %v", "127.0.0.1:6969", config.Address))

	c.Assert(config.IsClustered, qt.IsFalse,
		qt.Commentf("Config Test: Expected value %v, got %v", false, config.IsClustered))

	c.Assert(config.TestingEnabled, qt.IsFalse,
		qt.Commentf("Config Test: Expected value %v, got %v", false, config.TestingEnabled))

	c.Assert("", qt.Equals, config.MongodbURI,
		qt.Commentf("Config Test: Expected value %v, got %v", "(empty string)", config.MongodbURI))
}

func TestConfigInitialization_BooleanParsingError(t *testing.T) {

	c := qt.New(t)

	// Set environment variables for testing with invalid boolean values
	os.Setenv("IS_CLUSTERED", "invalid_boolean")
	os.Setenv("TESTING_ENABLED", "invalid_boolean")

	// Perform initialization
	config.Reset()

	// Assert the default values for boolean variables when parsing error occurs
	c.Assert(config.IsClustered, qt.IsFalse,
		qt.Commentf("Config Test: Expected value %v, got %v", false, config.IsClustered))

	c.Assert(config.TestingEnabled, qt.IsFalse,
		qt.Commentf("Config Test: Expected value %v, got %v", false, config.TestingEnabled))
}

func TestConfigInitialization_MissingMongoService(t *testing.T) {

	c := qt.New(t)

	// Set environment variables for testing with missing MONGO_SERVICE
	os.Setenv("MONGO_USERNAME", "test_user")
	os.Setenv("MONGO_PASSWORD", "test_password")
	os.Setenv("MONGO_SERVICE", "")

	// Perform initialization
	config.Reset()

	// Assert the default value for MongodbURI when MONGO_SERVICE is missing
	c.Assert("", qt.Equals, config.MongodbURI,
		qt.Commentf("Config Test: Expected value %v, got %v", "(empty string)", config.MongodbURI))
}
