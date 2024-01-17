package main

import (
	"github.com/csl991005/Atlantis/model"
	"github.com/csl991005/Atlantis/routes"
	"github.com/csl991005/Atlantis/utils"
)

func main() {
	utils.InitConfig()
	model.InitDb()
	routes.InitRouter()
}