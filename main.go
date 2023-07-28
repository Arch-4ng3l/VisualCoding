package main

import "github.com/Arch-4ng3l/VisualCoding/api"

func main() {
	server := api.NewAPIServer(":3000")
	server.Init()

}
