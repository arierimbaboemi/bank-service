package config

import (
	"github.com/arierimbaboemi/bank-service/domain"

	logger "github.com/arierimbaboemi/bank-lib-service/config"
	"github.com/jmoiron/sqlx"
)

/*
 * Implemtasi database dengan config dari .yaml
 */
func NewDBConnectionYAML() (*sqlx.DB, error) {
	config, err := domain.GetConfig()
	if err != nil {
		logger.GetLog().Fatal().Err(err).Msg("Failed to get config")
	}

	db, err := sqlx.Connect("mysql", config.GetDatabaseConfig())
	if err != nil {
		logger.GetLog().Fatal().Err(err).Msg("Failed to connect database")
	} else {
		logger.GetLog().Info().Msg("Database connected")
	}

	return db, nil
}

/*
 * Use database config from .env
 */
func NewDBConnectionENV() (*sqlx.DB, error) {
	config, err := domain.GetConfig()
	if err != nil {
		logger.GetLog().Fatal().Err(err).Msg("Failed to get config")
	}

	db, err := sqlx.Connect("mysql", config.GetDatabaseENVConfig())
	if err != nil {
		logger.GetLog().Fatal().Err(err).Msg("Failed to connect database")
	} else {
		logger.GetLog().Info().Msg("Database connected")
	}

	return db, nil
}
