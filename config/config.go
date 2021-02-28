//Package config is a configuration package. It can be used to store configuration data in multiple different pluggable backends
package config

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/a-berahman/gitpipe/utility/logger"
)

type config struct {
	Pipedrive Pipedrive `yaml:"PIPEDRIVE"`
	GitHub    GitHub    `yaml:"GITHUB"`
	MongoInfo MongoInfo `yaml:"MONGO_INFO"`
}

//Pipedrive presents Pipedrive API configuration
type Pipedrive struct {
	MainURL        string `yaml:"MAIN_URL"`
	AddActivityURL string `yaml:"ADD_ACTIVITY_URL"`
	TOKEN          string `yaml:"TOKEN"`
	GetActivityURL string `yaml:"GET_ACTIVITY_URL"`
}

//GitHub presents GitHub API configuration
type GitHub struct {
	MainURL  string `yaml:"MAIN_URL"`
	GistURL  string `yaml:"GIST_URL"`
	TOKEN    string `yaml:"TOKEN"`
	Username string `yaml:"USERNAME"`
}

// MongoInfo presents mongo db configuration
type MongoInfo struct {
	URL    string `yaml:"URL"`
	DBName string `yaml:"DB_NAME"`
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
func LoadConfig(configPath string) *DB {
	CFG.Pipedrive = Pipedrive{MainURL: os.Getenv("PIPEDRIVE_MAIN_URL"), AddActivityURL: os.Getenv("PIPEDRIVE_ADD_ACTIVITY_URL"), TOKEN: os.Getenv("PIPEDRIVE_TOKEN"), GetActivityURL: os.Getenv("PIPEDRIVE_GET_ACTIVITY_URL")}
	CFG.GitHub = GitHub{MainURL: os.Getenv("GITHUB_MAIN_URL"), GistURL: os.Getenv("GITHUB_GIST_URL"), TOKEN: os.Getenv("GITHUB_TOKEN"), Username: os.Getenv("GITHUB_USERNAME")}
	CFG.MongoInfo = MongoInfo{URL: os.Getenv("MONGO_URL"), DBName: os.Getenv("MONGO_DB_NAME")}

	logger.Initialize()
	dbInstance := getDBInstance()

	return dbInstance
}
