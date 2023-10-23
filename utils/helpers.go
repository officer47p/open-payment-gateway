package utils

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"open-payment-gateway/types"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/params"
	"github.com/joho/godotenv"
)

func getEnvVarOnlyIfExists(n string) string {
	v := os.Getenv(n)
	if v == "" {
		log.Fatalf("key %s does not exist in the environment variables", n)
	}

	return v
}

func LoadEnvVariableFile(path string) (types.EnvVariables, error) {
	env := types.EnvVariables{}

	err := godotenv.Load(path)
	if err != nil {
		return env, err
	}

	dbPort, err := strconv.ParseInt(getEnvVarOnlyIfExists("DB_PORT"), 10, 64)
	if err != nil {
		return env, errors.New("could not parse DB_PORT env variable")
	}
	env.DBPort = dbPort
	env.DBUrl = getEnvVarOnlyIfExists("DB_URL")
	env.DBName = getEnvVarOnlyIfExists("DB_NAME")
	env.DBUser = getEnvVarOnlyIfExists("DB_USER")
	env.DBPassword = getEnvVarOnlyIfExists("DB_PASSWORD")
	env.ProviderUrl = getEnvVarOnlyIfExists("PROVIDER_URL")
	env.NatsUrl = getEnvVarOnlyIfExists("NATS_URL")

	return env, nil
}

func WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func EtherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func StringToBigInt(s string) (*big.Int, bool) {
	n := big.Int{}
	return n.SetString(s, 10)
}
