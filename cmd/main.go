package main

import (
	"github.com/IamStubborN/test/app"
	_ "github.com/lib/pq"
)

func main() {
	app.NewApp().Run()
}
