package utils

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/params"
	"github.com/joho/godotenv"
)

type envVariables struct {
	DBUrl               string
	DBPort              int64
	DBName              string
	DBUser              string
	DBPassword          string
	ChainID             int64
	Decimals            int64
	NetworkName         string
	NetworkCurrency     string
	StartingBlockNumber int64
	ProviderUrl         string
	Environment         string // This can be prod or dev
}

func (e *envVariables) getOnlyIfExists(n string) string {
	v := os.Getenv(n)
	if v == "" {
		log.Fatalf("key %s does not exist in the environment variables", n)
	}

	return v
}

func isEnvironmentValid(e string) bool {
	return e == "prod" || e == "dev"
}

func LoadEnvVariableFile(path string) (envVariables, error) {
	env := envVariables{}

	err := godotenv.Load(path)
	if err != nil {
		return env, err
	}

	env.DBUrl = env.getOnlyIfExists("DB_URL")

	dbPort, err := strconv.ParseInt(env.getOnlyIfExists("DB_PORT"), 10, 64)
	if err != nil {
		return env, errors.New("could not parse DB_PORT env variable")
	}
	env.DBPort = dbPort

	env.DBName = env.getOnlyIfExists("DB_NAME")
	env.DBUser = env.getOnlyIfExists("DB_USER")
	env.DBPassword = env.getOnlyIfExists("DB_PASSWORD")

	chainId, err := strconv.ParseInt(env.getOnlyIfExists("CHAIN_ID"), 10, 64)
	if err != nil {
		return env, errors.New("could not parse CHAIN_ID env variable")
	}
	env.ChainID = chainId

	decimals, err := strconv.ParseInt(env.getOnlyIfExists("DECIMALS"), 10, 64)
	if err != nil {
		log.Fatal("Could not parse DECIMALS env variable")
	}
	env.Decimals = decimals

	env.ProviderUrl = env.getOnlyIfExists("PROVIDER_URL")
	env.NetworkName = env.getOnlyIfExists("NETWORK_NAME")
	env.NetworkCurrency = env.getOnlyIfExists("NETWORK_CURRENCY")

	startingBlockNumber, err := strconv.ParseInt(env.getOnlyIfExists("STARTING_BLOCK_NUMBER"), 10, 64)
	if err != nil {
		log.Fatal("Could not parse STARTING_BLOCK_NUMBER env variable")
	}
	env.StartingBlockNumber = startingBlockNumber

	e := env.getOnlyIfExists("ENVIRONMENT")
	fmt.Println(e)
	if !isEnvironmentValid(e) {
		log.Fatal("Invalid environment. Environment should be \"prod\" or \"dev\"")
	}
	env.Environment = e

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
