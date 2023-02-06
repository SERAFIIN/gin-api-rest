package main

import (
	"github.com/SERAFIIN/gin-api-rest/database"
	"github.com/SERAFIIN/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
