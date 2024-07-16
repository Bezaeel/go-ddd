package tests

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS orders (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    shipment_number INTEGER,
	    cargo_id INTEGER,
	    is_shipped BOOLEAN,
	    created_at DATETIME
	);

	CREATE TABLE IF NOT EXISTS order_line_items (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    product_id INTEGER,
	    seller_id INTEGER,
	    order_id INTEGER,
	    FOREIGN KEY (order_id) REFERENCES orders(id)
	);
	`

	// dropTableSQL := `
	// DROP TABLE orders;

	// DROP TABLE order_line_items;
	// `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return db
}

func ConvertToGormDB(sqlDB *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize GORM with existing sql.DB: %w", err)
	}
	return gormDB, nil
}

func TearDown(db *sql.DB) {

	dropTableSQL := `
	DROP TABLE orders;

	DROP TABLE order_line_items;
	`
	_, err := db.Exec(dropTableSQL)
	if err != nil {
		log.Fatalf("Failed to drop table: %v", err)
	}
}
