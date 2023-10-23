package types

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
