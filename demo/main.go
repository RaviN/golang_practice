package main

import (
	"net/http"

	"github.com/RaviN/golang_practice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
