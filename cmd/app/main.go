package main

import "go-jwt/internal/app"

const configsDir = "configs"

func main() {
	a := app.App{}

	a.Initialize(configsDir)
	a.Run()
}
