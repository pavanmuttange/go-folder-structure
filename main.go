package main

import (
	"github.com/pavanmuttange/go-folder-structure/pkg/database"
	"github.com/pavanmuttange/go-folder-structure/pkg/routes"
)

func main() {
	db := database.InitDB()
	r := routes.SetupRouter(db)
	r.Run(":8000")
}
