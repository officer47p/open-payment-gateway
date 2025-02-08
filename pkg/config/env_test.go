package config

// import (
// 	"fmt"
// 	"os"
// 	"testing"
// )

// const testDotenvFileName = "./.env.test"
// const dotEnvFileContent = "PROVIDER_URL=\"https://goerli.infura.io/v3/someToken\"\n" +
// 	"NATS_URL=\"nats://127.0.0.1:4222\"\n" +
// 	"DB_URL=\"localhost\"\n" +
// 	"DB_PORT=5432\n" +
// 	"DB_NAME=\"open-payment-gateway\"\n" +
// 	"DB_USER=\"postgres\"\n" +
// 	"DB_PASSWORD=\"postgres\"\n"

// func TestLoadEnvVariableFile(t *testing.T) {
// 	defer func() {
// 		err := os.Remove(testDotenvFileName)
// 		if err != nil {
// 			t.Errorf("Could not remove the test .env file: %v", err)
// 		}
// 	}()

// 	// Create the test .env file
// 	file, err := os.Create(testDotenvFileName)
// 	if err != nil {
// 		t.Fatalf("Could not create the test .env file with the path of: %s: %v", testDotenvFileName, err)
// 	}
// 	defer file.Close()

// 	_, err = file.WriteString(dotEnvFileContent)
// 	if err != nil {
// 		t.Fatalf("Could not write the dotenv file content to the path: %s: %v", testDotenvFileName, err)
// 	}

// 	// Load environment variables
// 	env, err := config.LoadEnvVariableFile(testDotenvFileName)
// 	if err != nil {
// 		t.Fatalf("LoadEnvVariableFile failed with the error: %s", err)
// 	}

// 	// Check if environment variables match
// 	checkEnvVar := func(name, expected, actual string) {
// 		if expected != actual {
// 			t.Errorf("Environment variable %s differs from what was written.\nExpected: %s\nActual: %s", name, expected, actual)
// 		}
// 	}

// 	checkEnvVar("PROVIDER_URL", "https://goerli.infura.io/v3/someToken", env.ProviderUrl)
// 	checkEnvVar("NATS_URL", "nats://127.0.0.1:4222", env.NatsUrl)
// 	checkEnvVar("DB_URL", "localhost", env.DBUrl)
// 	checkEnvVar("DB_PORT", "5432", fmt.Sprintf("%d", env.DBPort))
// 	checkEnvVar("DB_NAME", "open-payment-gateway", env.DBName)
// 	checkEnvVar("DB_USER", "postgres", env.DBUser)
// 	checkEnvVar("DB_PASSWORD", "postgres", env.DBPassword)
// }
