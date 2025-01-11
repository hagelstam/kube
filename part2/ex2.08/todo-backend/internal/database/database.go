package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Service interface {
	GetTodos() ([]Todo, error)
	CreateTodo(Todo) error
	Close() error
}

type service struct {
	db *sql.DB
}

var (
	database   = os.Getenv("POSTGRES_DB")
	password   = os.Getenv("POSTGRES_PASSWORD")
	user       = os.Getenv("POSTGRES_USER")
	host       = os.Getenv("POSTGRES_HOST")
	dbInstance *service
)

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable&search_path=public", user, password, host, database)
	log.Printf("Connecting to database: %s", connStr)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        text TEXT NOT NULL
    )`
	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatalf("Failed to create todos table: %v", err)
	}

	dbInstance = &service{db: db}
	return dbInstance
}

func (s *service) GetTodos() ([]Todo, error) {
	rows, err := s.db.Query("SELECT id, text FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Text)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *service) CreateTodo(todo Todo) error {
	_, err := s.db.Exec("INSERT INTO todos (text) VALUES ($1)", todo.Text)
	return err
}

func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}
