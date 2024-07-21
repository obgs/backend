package cmd

import (
	"fmt"

	"github.com/obgs/backend/internal/ent"
)

func createEntClient() (client *ent.Client, err error) {
	options := []ent.Option{}
	if config.EntDebug {
		options = append(options, ent.Debug())
	}

	return ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBAddress, config.DBPort, config.DBUser, config.DBPass, config.DBName, config.DBSSLMode), options...)
}
