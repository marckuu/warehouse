package main

import "Warehouse/server"

func main() {
	myServer := server.NewServerManager()

	myServer.StartServer()
}
