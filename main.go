package main

import (
	"final_project/database"
	"final_project/routes"
)

func main() {
	PORT := ":8080"
	database.SetupDB()
	routes.Serve().Run(PORT)

}
