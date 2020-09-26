package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	DBHosts             []string
	DBKeyspace          string
	DBCreateKeyspace    bool
	DBReplicationFactor int
	DBRunMigrations     bool
	Port                string
	HTTPSEnabled        bool
	TLSCert             string
	TLSKey              string
}

func LoadConfig(scope string) *Configuration {
	config := &Configuration{}
	setDefaults()
	switch scope {
	case "prod":
		loadConfigFile("config_prod")
		break
	case "dev":
		loadConfigFile("config_dev")
		break
	default:
		panic(fmt.Errorf("Scope not valid: %s. Must be either 'prod' or 'dev'.\n", scope))
	}
	setConfig(config)
	return config
}

func setDefaults() {
	viper.SetDefault("db_create_keyspace", false)
	viper.SetDefault("db_run_migrations", false)
	viper.SetDefault("https_enabled", false)
}

func loadConfigFile(configFilename string) {
	viper.SetConfigName(configFilename)
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func setConfig(config *Configuration) {
	checkMandatoryFields()

	config.DBHosts = viper.GetStringSlice("db_hosts")
	config.DBKeyspace = viper.GetString("db_keyspace")
	config.DBCreateKeyspace = viper.GetBool("db_create_keyspace")
	config.DBRunMigrations = viper.GetBool("db_run_migrations")
	config.Port = viper.GetString("port")
	config.HTTPSEnabled = viper.GetBool("https_enabled")

	if config.DBCreateKeyspace {
		config.DBReplicationFactor = viper.GetInt("db_replication_factor")
	}
	if config.HTTPSEnabled {
		config.TLSCert = viper.GetString("tls_cert")
		config.TLSKey = viper.GetString("tls_key")
	}
}

func checkMandatoryFields() {
	if !viper.IsSet("db_hosts") || !viper.IsSet("db_keyspace") || !viper.IsSet("port") {
		panic(fmt.Errorf("mandatory fields not set"))
	}
}
