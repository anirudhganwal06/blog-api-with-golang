package main

import (
	"fmt"

	"github.com/anirudhganwal06/blog-api-with-golang/app"
	"github.com/anirudhganwal06/blog-api-with-golang/config"
)

func main() {
	MongoURI := config.MongoURI
	fmt.Println("Starting application...")
	app := &app.App{}
	app.Initialize(MongoURI)
	var port string = "5000"
	fmt.Println(`Server running @` + port)
	app.Run(":" + port)
}
