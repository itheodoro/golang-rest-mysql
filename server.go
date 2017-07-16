package main

import (
	"github.com/itheodoro/rest-mysql/routes"
)

func main() {
	r := routes.InitRoutes()
	r.Run(":8080")
}
