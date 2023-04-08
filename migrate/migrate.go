package main

import (
	"github.com/atorahmad/go/crud/initializers"
	"github.com/atorahmad/go/crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.KonekDB()
}

func main() {
	initializers.Db.AutoMigrate(&models.Books{})
}
