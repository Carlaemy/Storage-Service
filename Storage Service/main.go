package main

import "./API/App"

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
