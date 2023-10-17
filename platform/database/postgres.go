package database

import (
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"github.com/sekke276/greendeco.git/pkg/configs"
)

type DB struct {
	*sqlx.DB
}

var defaultDB = &DB{}

// Connect postgresql
// This support to call once
func (db *DB) connectPostgresql() error {
	cfg := configs.AppConfig()
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode)
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

func ConnectDB() error {
	return defaultDB.connectPostgresql()
}

func (db *DB) loadSQLFile(sqlFile string) error {
	file, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		tx.Rollback()
	}()

	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := tx.Exec(q); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (db *DB) Migrate() error {
	if err := db.loadSQLFile("migrator/migrations/1_initialize_schema.up.sql"); err != nil {
		println(err.Error())
		return err
	}

	return nil
}
