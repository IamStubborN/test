package main

import (
	"github.com/IamStubborN/test/app"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	app.NewApp().Run()
}
