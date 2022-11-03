package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Datebase) MigrateDB() error {
	fmt.Println("migrating our database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)

	if err != nil {
		fmt.Println(err)
		return err

	}

	if err := m.Up(); err != nil {
		return fmt.Errorf("could not run UP migrations", err)
	}

	m.Up()

	fmt.Println("successfully migrated database")

	return nil
}
