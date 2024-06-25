package config

import (
	"errors"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var Config ConfigValues

// InitConfig set the value of config values
func InitConfig() {
	setDefaults()

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal(errors.New("config could not be loaded"))
	}
}

func setDefaults() {
	// Database
	viper.SetDefault("database.host", "ws://localhost:9999/rpc")
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "root")
	viper.SetDefault("database.namespace", "test")
	viper.SetDefault("database.database", "test")

	//Default
	viper.SetDefault("default.pagesize", 25)
	viper.SetDefault("default.botuseremail", "aibot@jmk21.com")

	//Security
	viper.SetDefault("security.bypassauth", "true")
	viper.SetDefault("security.jwtsigningkey", "local_signing_key_V4G7A1")
	viper.SetDefault("security.jwtaudience", "analyzer.net")
	viper.SetDefault("security.jwtissuer", "analyzer.net")

	//Server
	viper.SetDefault("server.enableplayground", "true")
	viper.SetDefault("server.serverport", 5000)
	viper.SetDefault("server.datadir", "/mnt/data")

	//Pub-Sub
	viper.SetDefault("pubsub.host", "nats://nats.nats.svc:4222")

	//Services
	viper.SetDefault("services.aihost", "http://localhost:7000/bot")
	viper.SetDefault("services.aihost", "http://ai.app.svc:7000/bot")
	//viper.SetDefault("services.ollamahost", "ollama.ollama.svc:11434")
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
}

type DefaultsConfig struct {
	PageSize     int
	BotUserEmail string
}

type SecurityConfig struct {
	BypassAuth    bool
	JwtSigningKey string
	JwtAudience   string
	JwtIssuer     string
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
