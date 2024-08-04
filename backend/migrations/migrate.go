package migrations

import (
	"log"

	"backend/infrastructure/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Could not create MySQL driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations/sql",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
