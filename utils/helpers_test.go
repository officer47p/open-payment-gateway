package utils

import (
	"fmt"
	"math/big"
	"os"
	"testing"
)

const testDotenvFileName = "./.env.test"
const dotEnvFileContent = "PROVIDER_URL=\"https://goerli.infura.io/v3/someToken\"\n" +
	"NATS_URL=\"nats://127.0.0.1:4222\"\n" +
	"DB_URL=\"localhost\"\n" +
	"DB_PORT=5432\n" +
	"DB_NAME=\"open-payment-gateway\"\n" +
	"DB_USER=\"postgres\"\n" +
	"DB_PASSWORD=\"postgres\"\n"

func TestLoadEnvVariableFile(t *testing.T) {
	defer func() {
		err := os.Remove(testDotenvFileName)
		if err != nil {
			t.Errorf("Could not remove the test .env file: %v", err)
		}
	}()

	// Create the test .env file
	file, err := os.Create(testDotenvFileName)
	if err != nil {
		t.Fatalf("Could not create the test .env file with the path of: %s: %v", testDotenvFileName, err)
	}
	defer file.Close()

	_, err = file.WriteString(dotEnvFileContent)
	if err != nil {
		t.Fatalf("Could not write the dotenv file content to the path: %s: %v", testDotenvFileName, err)
	}

	// Load environment variables
	env, err := LoadEnvVariableFile(testDotenvFileName)
	if err != nil {
		t.Fatalf("LoadEnvVariableFile failed with the error: %s", err)
	}

	// Check if environment variables match
	checkEnvVar := func(name, expected, actual string) {
		if expected != actual {
			t.Errorf("Environment variable %s differs from what was written.\nExpected: %s\nActual: %s", name, expected, actual)
		}
	}

	checkEnvVar("PROVIDER_URL", "https://goerli.infura.io/v3/someToken", env.ProviderUrl)
	checkEnvVar("NATS_URL", "nats://127.0.0.1:4222", env.NatsUrl)
	checkEnvVar("DB_URL", "localhost", env.DBUrl)
	checkEnvVar("DB_PORT", "5432", fmt.Sprintf("%d", env.DBPort))
	checkEnvVar("DB_NAME", "open-payment-gateway", env.DBName)
	checkEnvVar("DB_USER", "postgres", env.DBUser)
	checkEnvVar("DB_PASSWORD", "postgres", env.DBPassword)
}

// Compare two big.Float values with a tolerance for equality
func FloatsApproximatelyEqual(a, b *big.Float, tolerance float64) bool {
	diff := new(big.Float).Abs(new(big.Float).Sub(a, b))
	return diff.Cmp(big.NewFloat(tolerance)) <= 0
}

func TestWeiToEther(t *testing.T) {
	// Test cases with different values
	testCases := []struct {
		weiStr    string
		ethStr    string
		tolerance float64
	}{
		{"123000000000000000", "0.123", 0.0000000000000000001}, // Adjust tolerance as needed
		{"1234567890000000000", "1.23456789", 0.0000000000000000001},
		{"1234567897463527854", "1.234567897463527854", 0.0000000000000000001},
		{"100000000000000000000", "100", 0.0000000000000000001},
	}

	for _, tc := range testCases {
		wei, success := new(big.Int).SetString(tc.weiStr, 10)
		if !success {
			t.Errorf("Failed to convert test case input to big.Int: %s", tc.weiStr)
			continue
		}

		eth := WeiToEther(wei)
		ethStr := eth.Text('f', 18)

		expectedEth, success := new(big.Float).SetString(tc.ethStr)
		if !success {
			t.Errorf("Failed to convert test case expected value to big.Float: %s", tc.ethStr)
			continue
		}

		if !FloatsApproximatelyEqual(eth, expectedEth, tc.tolerance) {
			t.Errorf("Expected WeiToEther(%s) to be approximately %s, but got %s", tc.weiStr, tc.ethStr, ethStr)
		}
	}
}

func TestEtherToWei(t *testing.T) {
	// Test cases with different values
	testCases := []struct {
		ethStr string
		weiStr string
	}{
		{"0.123", "123000000000000000"},
		{"1.234567897462454854", "1234567897462454854"},
		{"100", "100000000000000000000"},
	}

	for _, tc := range testCases {
		eth, success := new(big.Float).SetString(tc.ethStr)
		if !success {
			t.Errorf("Failed to convert test case input to big.Float: %s", tc.ethStr)
			continue
		}

		wei := EtherToWei(eth)
		weiStr := wei.String()

		if weiStr != tc.weiStr {
			t.Errorf("Expected EtherToWei(%s) to be %s, but got %s", tc.ethStr, tc.weiStr, weiStr)
		}
	}
}

func TestStringToBigInt(t *testing.T) {
	// Test cases with different input strings and their expected results
	testCases := []struct {
		input       string
		expected    *big.Int
		expectError bool
	}{
		{"123", big.NewInt(123), false},
		{"0", big.NewInt(0), false},
		{"-456", big.NewInt(-456), false},
		{"4829376597832654983", big.NewInt(4829376597832654983), false},
		{"invalid", nil, true},
		{"", nil, true},
	}

	for _, tc := range testCases {
		result, ok := StringToBigInt(tc.input)

		if tc.expectError {
			if ok {
				t.Errorf("Expected not ok for input: %s", tc.input)
			}
		} else {
			if !ok {
				t.Errorf("Unexpected ok for input: %s - %v", tc.input, ok)
			} else if result.Cmp(tc.expected) != 0 {
				t.Errorf("Expected StringToBigInt(%s) to be %s, but got %s", tc.input, tc.expected.String(), result.String())
			}
		}
	}
}
