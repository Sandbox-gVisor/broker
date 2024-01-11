package storages

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

// PostgresStorage implements Storage interface
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

	//creating database
	initDatabase(store.ConnectionPool, store.Ctx)
	log.Println("Database initialized")
}

func initDatabase(pool *pgxpool.Pool, ctx context.Context) {
	// Creating table if it doesn't exist
	tag, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS messages (
		    id      bigserial,
		    message jsonb
		)
	`)

	log.Println("Table was successfully created! Tag: ", tag)
	if err != nil {
		log.Printf("Unable to create table: %v\n", err)
		os.Exit(1)
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
	_, err := store.ConnectionPool.Exec(store.Ctx, `INSERT INTO messages (message) VALUES (CAST ($1 AS jsonb))`, msg)
	if err != nil {
		log.Println("Couldn't insert into table messages!")
	} else {
		log.Println("Successfully inserted into messages!")
	}
}
