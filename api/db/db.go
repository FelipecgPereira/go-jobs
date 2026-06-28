package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "go_job.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	createTables()
	createIndexes()

}

func createTables() {
	// Table user
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		name TEXT NOT NULL,
		create_at DATETIME,
		update_at DATETIME
	);
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	// Table Customers
	createCustomersTable := `
	CREATE TABLE IF NOT EXISTS customers(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		phone TEXT,
		create_at DATETIME,
		update_at DATETIME
	);
	`
	_, err = DB.Exec(createCustomersTable)
	if err != nil {
		panic(err)
	}

	// Table Projects
	createProjectTable := `
	CREATE TABLE IF NOT EXISTS projects(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price DECIMAL (8,2),
		start_date DATETIME,
		end_date DATETIME,
		create_at DATETIME,
		update_at DATETIME
	);
	`
	_, err = DB.Exec(createProjectTable)
	if err != nil {
		panic(err)
	}

	// Table b2b
	createB2BTable := `
	CREATE TABLE IF NOT EXISTS b2b(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_id INTEGER,
		user_id INTEGER,
		project_id INTEGER,
		FOREIGN KEY(customer_id) REFERENCES customers(id)
		FOREIGN KEY(user_id) REFERENCES users(id)
		FOREIGN KEY(project_id) REFERENCES projects(id)
	);
	`
	_, err = DB.Exec(createB2BTable)
	if err != nil {
		panic(err)
	}

	// Initialize ID sequences to start at 10000
	initializeSequences()
}

func initializeSequences() {
	tables := []string{"users", "customers", "projects", "b2b"}

	for _, table := range tables {
		query := `INSERT OR IGNORE INTO sqlite_sequence (name, seq) VALUES (?, ?)`
		_, err := DB.Exec(query, table, 9999)
		if err != nil {
			panic(err)
		}
	}
}

func createIndexes() {
	indexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);`,
		`CREATE INDEX IF NOT EXISTS idx_customers_email ON customers(email);`,
		`CREATE INDEX IF NOT EXISTS idx_b2b_customer_id ON b2b(customer_id);`,
		`CREATE INDEX IF NOT EXISTS idx_b2b_user_id ON b2b(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_b2b_project_id ON b2b(project_id);`,
	}

	for _, indexSQL := range indexes {
		_, err := DB.Exec(indexSQL)
		if err != nil {
			panic(err)
		}
	}
}
