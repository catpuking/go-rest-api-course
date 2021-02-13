package main

import (
	"fmt"
)

//App - the struct which contains things like
//pointers to DB connections
type App struct{}


// Run - Sets up our application
func  (app *App) Run() error {
	fmt.Println("Setting up our APP")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting our REST API")
		fmt.Println(err)
	}
}