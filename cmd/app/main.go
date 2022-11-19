package main

import "grain-acceptance/internal/runner"

const configDir = "./config/"

func main() {
	runner.Start(configDir)
}
