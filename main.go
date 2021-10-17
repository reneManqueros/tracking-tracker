package main

import (
	"trackingtracker/cmd"
	"trackingtracker/internal/config"
)

func init() {
	config.Load()
}
func main() {
	cmd.Execute()
}
