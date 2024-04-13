package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

func newDatabase(ctx context.Context, applicationName string) model.PgxPool {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("could not parse config pgxpool, err: %v", err)
	}

	config.ConnConfig.RuntimeParams["application_name"] = applicationName

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("could not connection to db, err: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("could ping database, err: %v", err)
	}

	return db
}
