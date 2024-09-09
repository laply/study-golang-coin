package main

import (
	"log"

	app "github.com/laply/study-golang-coin/internal"
)

func main() {

	App := app.New()

	if err := App.Run(); err != nil {
		log.Fatal(err)
	}

}
