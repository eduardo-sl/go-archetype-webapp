package main

import "github.com/eduardo-sl/go-archetype-webapp/internal/api"

func main() {
	app := api.NewAPI()
	app.Run()
}
