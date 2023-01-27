package cmd

import (
	"log"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

var cfgFile string

const (
	keepAlivePingInterval            = 10 * time.Second
	queryCacheLruSize                = 1000
	automaticPersistedQueryCacheSize = 100
)

type Config struct {
	DBAddress               string `mapstructure:"DB_ADDRESS"`
	DBPort                  string `mapstructure:"DB_PORT"`
	DBName                  string `mapstructure:"DB_NAME"`
	DBUser                  string `mapstructure:"DB_USER"`
	DBPass                  string `mapstructure:"DB_PASS"`
	JWTSecret               string `mapstructure:"JWT_SECRET"`
	ServerHost              string `mapstructure:"SERVER_HOST"`
	OAuthGoogleClientID     string `mapstructure:"OAUTH_GOOGLE_CLIENT_ID"`
	OAuthGoogleClientSecret string `mapstructure:"OAUTH_GOOGLE_CLIENT_SECRET"`
	S3AccessKeyID           string `mapstructure:"S3_ACCESS_KEY_ID"`
	S3SecretAccessKey       string `mapstructure:"S3_SECRET_ACCESS_KEY"`
	S3Region                string `mapstructure:"S3_REGION"`
	S3Endpoint              string `mapstructure:"S3_ENDPOINT"`
	S3Bucket                string `mapstructure:"S3_BUCKET"`
	UsingMinio              bool   `mapstructure:"USING_MINIO"`
	EntDebug                bool   `mapstructure:"ENT_DEBUG"`
	IntrospectionEnabled    bool   `mapstructure:"INTROSPECTION_ENABLED"`
}

// https://github.com/spf13/viper/issues/761
func bindEnvs() {
	viper.BindEnv("DB_ADDRESS")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASS")
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("SERVER_HOST")
	viper.BindEnv("OAUTH_GOOGLE_CLIENT_ID")
	viper.BindEnv("OAUTH_GOOGLE_CLIENT_SECRET")
	viper.BindEnv("S3_ACCESS_KEY_ID")
	viper.BindEnv("S3_SECRET_ACCESS_KEY")
	viper.BindEnv("S3_REGION")
	viper.BindEnv("S3_ENDPOINT")
	viper.BindEnv("S3_BUCKET")
	viper.BindEnv("USING_MINIO")
	viper.BindEnv("ENT_DEBUG")
	viper.BindEnv("INTROSPECTION_ENABLED")
}

var config Config

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	bindEnvs()

	viper.SetConfigFile(cfgFile)

	configType := filepath.Ext(cfgFile)
	viper.SetConfigType(configType[1:])

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	log.Default().Printf("Using config file: %s", viper.ConfigFileUsed())
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
}
