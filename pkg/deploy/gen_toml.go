package deploy

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Master []MasterConfig `toml:"master"`
	Client []ClientConfig `toml:"client"`
}

type MasterConfig struct {
	IP     string `toml:"ip"`
	Port   int    `toml:"port"`
	User   string `toml:"user"`
	Passwd string `toml:"passwd"`
}

type ClientConfig struct {
	IP     string `toml:"ip"`
	Port   int    `toml:"port"`
	User   string `toml:"user"`
	Passwd string `toml:"passwd"`
}

func genToml() {
	// Define config
	config := Config{
		Master: []MasterConfig{
			{
				IP:     "192.168.1.1",
				Port:   22,
				User:   "corey",
				Passwd: "corey1996",
			},
		},
		Client: []ClientConfig{
			{
				IP:     "192.168.1.3",
				Port:   22,
				User:   "corey",
				Passwd: "corey1996",
			},
		},
	}

	// Encode config to TOML
	file, err := os.Create("config.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err := toml.NewEncoder(file).Encode(config); err != nil {
		panic(err)
	}

	fmt.Println("Config file generated successfully!")
}

func GetTomlConf(path string) *Config {
	// Decode config from TOML
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		panic(err)
	}
	return &conf
}
