package main

import "fmt"

// Run - is going to be responsible for
// instantiation and startup of our app
func Run() error {
	fmt.Println("starting our app")
	return nil
}
func main() {
	fmt.Println("Go REST API course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
