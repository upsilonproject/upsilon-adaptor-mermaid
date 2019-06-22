package main;

import "github.com/BurntSushi/toml"

type DatabaseConfig struct {
	User string;
	Pass string;
	Host string;
	Name string;
}

type NetworkConfig struct {
	Port int64;
}

type Config struct {
	Database DatabaseConfig;
	Network NetworkConfig;
}

func GetConfig() Config {
	var conf Config;

	if _, err := toml.DecodeFile("/etc/upsilon-adaptor-mermaid/config.toml", &conf); err != nil {
		panic(err);
	}

	return conf;
}

