package configurations

import "github.com/spf13/viper"

type Configuration struct {
	Environment string
	Mongo       MongoConfiguration
}

type MongoConfiguration struct {
	Server     string
	Database   string
	Collection string
}

func GetConfiguration() Configuration {
	configuration := Configuration{}

	viper.SetConfigName("configuration")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configurations")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
