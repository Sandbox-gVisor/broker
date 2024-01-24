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

// initDatabase creates table in database and sets trigger on created table, that
// notifies ws_server about changes
func initDatabase(pool *pgxpool.Pool, ctx context.Context) {
	// Creating table if it doesn't exist
	tag, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS messages (
		    id      bigserial,
		    message jsonb
		)
	
	`)

	if err != nil {
		log.Printf("Unable to create table: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("Table was successfully created! Tag: ", tag)
	}

	tag, err = pool.Exec(ctx, `
          CREATE OR REPLACE FUNCTION notify_ws_server() RETURNS TRIGGER AS $$
          BEGIN
			PERFORM pg_notify('update', row_to_json(NEW)::text);
			RETURN NULL;
		  END;
		  $$ LANGUAGE plpgsql;

		  CREATE TRIGGER messages_table_change
		  AFTER INSERT OR UPDATE OR DELETE ON messages
		  FOR EACH ROW EXECUTE FUNCTION notify_ws_server();
	`)

	if err != nil {
		log.Printf("Unable to create trigger: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("Trigger was successfully created! Tag: ", tag)
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
	}
}
