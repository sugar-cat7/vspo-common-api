package main

import (
	"log"
	"net/http"

	_ "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/channel"
	_ "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/clip"
	_ "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/song"
)

// @title VSPO Common API
// @version 1.0
// @description This is the API documentation for VSPO Common services.
// @BasePath /api/v1

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
