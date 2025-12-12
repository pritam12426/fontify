package main

import (
	"fmt"
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
