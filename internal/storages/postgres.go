package storages

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type PostgresStorage struct {
	ConnectionPool *pgxpool.Pool
	Ctx            context.Context
}

func (store *PostgresStorage) Init() {
	var err error

	store.Ctx = context.Background()
	store.ConnectionPool, err = pgxpool.New(store.Ctx, os.Getenv("POSTGRES_ADDR"))
	if err != nil {
		log.Printf("Unable to connect to database: %v\n\n", err)
		os.Exit(1)
	}

	err = store.ConnectionPool.Ping(store.Ctx)
	if err != nil {
		log.Println("Couldn't ping postgres!")
	}
}

func (store *PostgresStorage) Close() {
	err := store.ConnectionPool.Close
	if err != nil {
		log.Println("Couldn't close connection!")
		os.Exit(1)
	}
}

func (store *PostgresStorage) FlushStorage() {
	// ??
}

func (store *PostgresStorage) SaveMessage(msg string) {
	_, err := store.ConnectionPool.Exec(store.Ctx, `insert into `)
	if err != nil {
		log.Println("Couldn't start transaction!")
	}
}
