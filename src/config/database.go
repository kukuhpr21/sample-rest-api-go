package config

import (
	"database/sql"
	"time"
)

type DatabaseConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func NewDB(dbConfig DatabaseConfig) (*sql.DB, error) {

	// Db Config
	 mPassword := ""
	if len(dbConfig.Password) > 0 {
		mPassword = ":"+dbConfig.Password
	}
	
	mydb := dbConfig.Username +mPassword + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name
	db, err := sql.Open(dbConfig.Driver, mydb)

	if err == nil {
		// Db Pooling
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(20)
		db.SetConnMaxLifetime(60 * time.Minute)
		db.SetConnMaxIdleTime(10 * time.Minute)
	}

	return db, err
}
