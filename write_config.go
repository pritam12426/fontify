package main

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title string
	Owner struct {
		Name string
		Dob  time.Time
	}
	Database struct {
		Server        string
		Ports         []int
		ConnectionMax int `toml:"connection_max"`
		Enabled       bool
	}
}

func main() {
	// Create a new configuration instance
	config := Config{
		Title: "New TOML Config",
		Owner: struct {
			Name string
			Dob  time.Time
		}{
			Name: "Jane Doe",
			Dob:  time.Now(),
		},
		Database: struct {
			Server        string
			Ports         []int
			ConnectionMax int `toml:"connection_max"`
			Enabled       bool
		}{
			Server:        "192.168.2.1",
			Ports:         []int{9000, 9001},
			ConnectionMax: 10000,
			Enabled:       false,
		},
	}

	// Create a new file to write to
	file, err := os.Create("new_config.toml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode the configuration to the file
	if err := toml.NewEncoder(file).Encode(config); err != nil {
		fmt.Println("Error encoding TOML to file:", err)
		return
	}

	fmt.Println("Successfully wrote configuration to new_config.toml")
}
