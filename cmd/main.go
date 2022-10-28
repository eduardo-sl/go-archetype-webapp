package main

import "github.com/eduardo-sl/go-webapp-project/internal/api"

func main() {
	app := api.NewAPI()
	app.Run()
}
