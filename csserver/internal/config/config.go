package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var Config ConfigValues

// InitConfig set the value of config values
func InitConfig() {
	setDefaults()

	//---TODO: figure out where to store the config
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(errors.New("could not find home directory"))
	}

	// Search config in home directory with name ".csplanner" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".csplanner")

	// If a config file is found, read it in.
	if err = viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("csp")
	viper.AutomaticEnv() // read in environment variables that match

	// Config = readInConfig()
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal(errors.New("config could not be loaded"))
	}
}

func setDefaults() {
	// Database
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "root")
	viper.SetDefault("database.namespace", "test")
	viper.SetDefault("database.database", "test")
	viper.SetDefault("database.port", 9999)

	//Default
	viper.SetDefault("default.pagesize", 25)
	viper.SetDefault("default.botuseremail", "aibot@jmk21.com")

	//Security
	viper.SetDefault("security.bypassauth", "false")
	viper.SetDefault("security.jwtaudience", "csplanner.io")
	viper.SetDefault("security.jwtissuer", "csplanner.io")
	viper.SetDefault("security.keycloakurl", "http://keycloak-keycloakx-http.keycloak.svc/auth")
	viper.SetDefault("security.keycloakrealm", "csplanner")
	viper.SetDefault("security.keycloakadminuser", "csplanner-admin")
	viper.SetDefault("security.keycloakadminpass", "pass")
	viper.SetDefault("security.keycloakclientid", "web-login")
	viper.SetDefault("security.keycloakclientsecret", "pass")

	//Server
	viper.SetDefault("server.enableplayground", "true")
	viper.SetDefault("server.serverport", 5000)
	viper.SetDefault("server.datadir", "/mnt/data")

	//Pub-Sub
	viper.SetDefault("pubsub.host", "nats://nats.nats.svc:4222")

	//Services
	viper.SetDefault("services.aihost", "http://localhost:7000/bot")
	viper.SetDefault("services.ollamahost", "ollama.ollama.svc:11434")
}

// Config the parent config structure
type ConfigValues struct {
	Database DatabaseConfig
	Default  DefaultsConfig
	Security SecurityConfig
	Server   ServerConfig
	PubSub   PubSubConfig
	Services ServicesConfig
}

// DatabaseConfig configuration values specific to the database
type DatabaseConfig struct {
	Host      string
	User      string
	Password  string
	Namespace string
	Database  string
	Port      int
}

type DefaultsConfig struct {
	PageSize     int
	BotUserEmail string
}

type SecurityConfig struct {
	BypassAuth           bool
	JwtSigningKey        string
	JwtAudience          string
	JwtIssuer            string
	KeycloakURL          string
	KeycloakRealm        string
	KeycloakAdminUser    string
	KeycloakAdminPass    string
	KeycloakClientID     string
	KeycloakClientSecret string
}

type ServerConfig struct {
	EnablePlayground bool
	ServerPort       int
	DataDir          string
}

type PubSubConfig struct {
	Host string
}

type ServicesConfig struct {
	AIHost     string
	OllamaHost string
}
