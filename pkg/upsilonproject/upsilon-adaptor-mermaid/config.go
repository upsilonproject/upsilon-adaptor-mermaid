package upsilonAdaptorMermaid;

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database DatabaseConfig;
	Network NetworkConfig;
	IsLoaded bool;
}

type DatabaseConfig struct {
	User string;
	Pass string;
	Host string;
	Name string;
}

type NetworkConfig struct {
	Port int64;
}

func initConfig() {
	viper.SetEnvPrefix("UP_MERMAID")
	viper.SetDefault("port", 8080);
	viper.SetConfigName("upsilon-adaptor-mermaid");
	viper.AddConfigPath("/etc/upsilon-adaptor-mermaid");

	err := viper.ReadInConfig();

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Config not found: %s \n", err)
		} else {
			// Config file was found but another error was produced
			log.Printf("Fatal error config file: %s \n", err)
		}
	}
}

var conf Config;

func GetConfig() Config {
	if !conf.IsLoaded {
		log.Println("Getting config from disk");

		initConfig();

		if err := viper.Unmarshal(&conf); err != nil {
			log.Println("Could not read config: %s", err);
		}

		conf.IsLoaded = true;

		log.Printf("Config is: \n %+v", conf);
	} else {
		log.Println("Getting config from cache");
	}

	return conf;
}

