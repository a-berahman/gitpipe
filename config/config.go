//Package config is a configuration package. It can be used to store configuration data in multiple different pluggable backends
package config

import (
	"fmt"

	"strings"

	"gopkg.in/yaml.v2"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type config struct {
	Pipedrive Pipedrive `yaml:"PIPEDRIVE"`
	GitHub    GitHub    `yaml:"GITHUB"`
}

type Pipedrive struct {
	MainURL        string `yaml:"MAIN_URL"`
	AddActivityURL string `yaml:"ADD_ACTIVITY_URL"`
	TOKEN          string `yaml:"TOKEN"`
}

type GitHub struct {
	MainURL  string `yaml:"MAIN_URL"`
	GistURL  string `yaml:"GIST_URL"`
	TOKEN    string `yaml:"TOKEN"`
	Username string `yaml:"USERNAME"`
}

type SetupResult struct {
}

type unmarshaler struct{}

//Unmarshal converts bytes to human readable format
func (unmarshaler) Unmarshal(d []byte, v *map[string]map[string]string) error {
	err := yaml.Unmarshal(d, v)
	return err
}

//CFG is config instance
var CFG config

//LoadConfig loads and initializes config list
func LoadConfig(configPath string) {

	viper.SetEnvPrefix("GITPIPE")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigFile(configPath)
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.MergeInConfig()
	if err != nil {
		fmt.Println("Error in reading config")
		panic(err)
	}

	err = viper.Unmarshal(&CFG, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
	})

	if err != nil {
		fmt.Println("Error in un-marshaling config")
		panic(err)
	}

}
