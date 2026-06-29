package mysql

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"rent-manager-backend/internal/adapter/config"

	"github.com/Masterminds/squirrel"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/go-sql-driver/mysql"
)

// migrationsFS embeds the migrations folder.
//
//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	*sql.DB
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

func New(ctx context.Context, config *config.DB) (*DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	migrationURL := fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}

	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

	return &DB{
		DB:           db,
		QueryBuilder: &builder,
		url:          migrationURL,
	}, nil
}

func (db *DB) Migrate() error {
	driver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, db.url)
	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func (db *DB) ErrorCode(err error) string {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return fmt.Sprintf("%d", mysqlErr.Number)
	}

	return ""
}

func (db *DB) Close() {
	_ = db.DB.Close()
}
