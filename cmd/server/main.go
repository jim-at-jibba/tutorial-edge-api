package main

import (
	"fmt"

	"github.com/jim-at-jibba/comments-api/internal/comment"
	"github.com/jim-at-jibba/comments-api/internal/db"
	transportHttp "github.com/jim-at-jibba/comments-api/internal/transport/http"
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
	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}
func main() {
	fmt.Println("Go REST API course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
