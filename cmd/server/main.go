package main

import (
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully conected and pinged database")
	// cmtService := comment.NewService(db)
	// fmt.Println(cmtService.GetComment(
	// 	context.Background(),
	// 	"2800e3a2-c904-4180-abe9-d75d55fdff4b",
	// ))
	return nil
}
func main() {
	fmt.Println("Go REST API course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
