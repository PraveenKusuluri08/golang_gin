package main

import "github.com/Praveenkusuri08/routes"

func main() {
	r := routes.TodoGetRoute()

	r.Run(":8000")
}
