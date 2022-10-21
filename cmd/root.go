package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

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
}

var config Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "backend",
	Short: "OBGS backend",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".backend.yaml", "config file (default is .backend.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetConfigType("yaml")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	log.Default().Printf("Using config file: %s", viper.ConfigFileUsed())
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
}
