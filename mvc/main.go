package main

import (
	"github.com/samykaplan/test-microservices/mvc/app"
)

func main() {
	app.StartApp()

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatalln(err)
	// }
}
