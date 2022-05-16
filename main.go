package main

import (
	"blogo/model"
	"blogo/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
