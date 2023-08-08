package database

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"github.com/sekke276/greendeco.git/pkg/configs"
)

type DB struct {
	*sqlx.DB
}

var defaultDB = &DB{}

func (db *DB) ConnectPostgresql(cfg *configs.Config) error {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port)
	db.DB = sqlx.MustOpen("pgx", dns)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		defer db.Close()
		return fmt.Errorf("Can not send ping to database, %w", err)
	}
	return nil
}

func GetDB() *DB {
	return defaultDB
}

func ConnectDB(cfg *configs.Config) error {
	return defaultDB.ConnectPostgresql(cfg)
}
