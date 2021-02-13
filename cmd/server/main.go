package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/catpuking/go-rest-api-course/internal/transport/http"
)

//App - the struct which contains things like
//pointers to DB connections
type App struct{}


// Run - Sets up our application
func  (app *App) Run() error {
	fmt.Println("Setting up our APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

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