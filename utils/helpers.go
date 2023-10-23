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

type EnvVariables struct {
	DBUrl       string
	DBPort      int64
	DBName      string
	DBUser      string
	DBPassword  string
	ProviderUrl string
	NatsUrl     string
	Environment string // This can be prod or dev
}

func (e *EnvVariables) getOnlyIfExists(n string) string {
	v := os.Getenv(n)
	if v == "" {
		log.Fatalf("key %s does not exist in the environment variables", n)
	}

	return v
}

func LoadEnvVariableFile(path string) (EnvVariables, error) {
	env := EnvVariables{}

	err := godotenv.Load(path)
	if err != nil {
		return env, err
	}

	dbPort, err := strconv.ParseInt(env.getOnlyIfExists("DB_PORT"), 10, 64)
	if err != nil {
		return env, errors.New("could not parse DB_PORT env variable")
	}
	env.DBPort = dbPort
	env.DBUrl = env.getOnlyIfExists("DB_URL")
	env.DBName = env.getOnlyIfExists("DB_NAME")
	env.DBUser = env.getOnlyIfExists("DB_USER")
	env.DBPassword = env.getOnlyIfExists("DB_PASSWORD")
	env.ProviderUrl = env.getOnlyIfExists("PROVIDER_URL")
	env.NatsUrl = env.getOnlyIfExists("NATS_URL")

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
