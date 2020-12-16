package main

import (
	"jmolboy/week04/internal/dao"
	"jmolboy/week04/internal/routes"
)

func main()  {
	dao.Init()

	routes.Start()
}

