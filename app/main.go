package main

import (
	"fmt"

	"github.com/lfsmariz/manager-system-backend/app/routes"
)

func main() {

	s := routes.SetupRouter()

	s.Run()

	fmt.Println("initial")
}
