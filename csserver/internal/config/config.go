package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
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

	//---TODO: make this less embarrassing
	envPaths := []string{
		"../../../../../secrets/.env.local",
		"../../../../secrets/.env.local",
		"../../../secrets/.env.local",
		"../../secrets/.env.local",
		"../secrets/.env.local",
		"./secrets/.env.local",
	}

	for _, p := range envPaths {
		fp, err := filepath.Abs(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot find .env.local.  Moving on...")
		} else {
			err = godotenv.Load(fp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Cannot load .env.local in %s.  Moving on...\n", fp)
			} else {
				//---if we load an env file...no need to keep looking
				fmt.Fprintf(os.Stderr, ".env.local loaded from %s.  test setttings applied\n", fp)
				break
			}
		}
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
	viper.SetDefault("masterdb.host", "postgres-postgresql.database.svc")
	viper.SetDefault("masterdb.user", "postgres")
	viper.SetDefault("masterdb.password", "postgres")
	viper.SetDefault("masterdb.namespace", "")
	viper.SetDefault("masterdb.database", "cssaas")
	viper.SetDefault("masterdb.port", 5432)

	// Database
	viper.SetDefault("database.host", "postgres-postgresql.database.svc")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "postgres")
	viper.SetDefault("database.namespace", "")
	viper.SetDefault("database.database", "csplanner")
	viper.SetDefault("database.port", 5432)

	//Default
	viper.SetDefault("default.pagesize", 25)
	viper.SetDefault("default.botuseremail", "aibot@jmk21.com")
	viper.SetDefault("default.botpassword", "pass")

	//Security
	viper.SetDefault("security.bypassauth", "false")
	viper.SetDefault("security.jwtaudience", "csplanner.io")
	viper.SetDefault("security.jwtissuer", "csplanner.io")
	viper.SetDefault("security.keycloakurl", "http://keycloak-keycloakx-http.keycloak.svc/auth")
	viper.SetDefault("security.keycloakrealm", "csplanner")
	viper.SetDefault("security.keycloakadminuser", "csplanner-admin")
	viper.SetDefault("security.keycloakadminpass", "pass")
	viper.SetDefault("security.keycloakclientid", "web-client")
	viper.SetDefault("security.keycloakclientsecret", "pass")
	viper.SetDefault("security.keycloakrealmtemplatepath", "")
	viper.SetDefault("security.keycloakmasterrealmname", "master")

	viper.SetDefault("security.keycloakprovisionerclientid", "cs-provisioner")
	viper.SetDefault("security.keycloakprovisionerclientsecret", "pass")

	//Server
	viper.SetDefault("server.enableplayground", "true")
	viper.SetDefault("server.serverport", 5000)
	viper.SetDefault("server.datadir", "/mnt/data")

	//Pub-Sub
	viper.SetDefault("pubsub.host", "nats://nats.nats.svc:4222")
	viper.SetDefault("pubsub.subjectformat", "%s.%s.%s.%s")
	viper.SetDefault("pubsub.streamname", "CSPLANNER")
	viper.SetDefault("pubsub.name", "csplanner-nats")

	//Services
	viper.SetDefault("services.aihost", "http://localhost:7000/bot")
	viper.SetDefault("services.ollamahost", "ollama.ollama.svc:11434")
}

// Config the parent config structure
type ConfigValues struct {
	MasterDB DatabaseConfig
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
	BotPassword  string
}

type SecurityConfig struct {
	BypassAuth                      bool
	JwtSigningKey                   string
	JwtAudience                     string
	JwtIssuer                       string
	KeycloakURL                     string
	KeycloakMasterRealmName         string
	KeycloakRealm                   string
	KeycloakAdminUser               string
	KeycloakAdminPass               string
	KeycloakClientID                string
	KeycloakClientSecret            string
	KeycloakRealmTemplatePath       string
	KeycloakProvisionerClientID     string
	KeycloakProvisionerClientSecret string
}

type ServerConfig struct {
	EnablePlayground bool
	ServerPort       int
	DataDir          string
}

type PubSubConfig struct {
	Host          string
	SubjectFormat string
	StreamName    string
	Name          string
}

type CMSConfig struct {
	OrgID   string
	SpaceID string
	PAT     string
}

type ServicesConfig struct {
	AIHost     string
	OllamaHost string
}
