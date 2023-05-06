package main

import "echo_rest_api/routers"

func main() {
	e := routers.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
