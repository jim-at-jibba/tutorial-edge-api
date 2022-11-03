package main

import (
	"context"
	"fmt"

	"github.com/jim-at-jibba/comments-api/internal/db"
)

// Run - is going to be responsible for
// instantiation and startup of our app
func Run() error {
	fmt.Println("starting our app")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	// This is redundent as the Connect methond pings but
	// for now it means db is used
	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("successfully conected and pinged database")
	return nil
}
func main() {
	fmt.Println("Go REST API course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
