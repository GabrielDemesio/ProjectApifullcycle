package main

import "FULLCYCLE/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

}
