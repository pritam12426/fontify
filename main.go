package main

import (
// 	// "github.com/BurntSushi/toml"
// 	// "github.com/spf13/cobra"
// 	// "log"
// 	// "net/http"
"fmt"
)

func main() {
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println("Error decoding TOML file:", err)
		return
	}

	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Owner: %s\n", config.Owner.Name)
	fmt.Printf("Owner's DOB: %s\n", config.Owner.Dob)
	fmt.Printf("Database server: %s\n", config.Database.Server)
	fmt.Printf("Database ports: %v\n", config.Database.Ports)
	fmt.Printf("Database connection max: %d\n", config.Database.ConnectionMax)
	fmt.Printf("Database enabled: %t\n", config.Database.Enabled)
}
