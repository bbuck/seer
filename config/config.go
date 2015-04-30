package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/tree-server/trees/errors"
	"github.com/tree-server/trees/log"
)

type Database struct {
	Host string `toml:"host"`
	//Auth bool
	//User string
	//Pass string
}

type Config struct {
	DB       Database `toml:"database"`
	RootNode string   `toml:"root_node_id"`
}

const configFileName = "Trees.toml"

var (
	loaded       = false
	loadedConfig Config
	logger       log.Logger
)

func init() {
	logger = log.Make("config", log.GetFileTarget(":stdout:"))
}

func LoadOrCreate() {
	err := load()

	if os.IsNotExist(err) {
		f, err := os.Create(configFileName)
		if err != nil {
			logger.Fatal(err, errors.ErrCreateConfigFailed)
		}
		defer f.Close()

		_, err = f.Write([]byte(defaultConfigToml))
		if err != nil {
			logger.Fatal(err, errors.ErrCreateConfigFailed)
		}

		logger.Log(log.LogInfo, "Created default configuartion.")
	} else {
		logger.Log(log.LogInfo, "Configuration already exists, doing nothing.")
	}
}

func load() error {
	loaded = true
	loadedConfig = withDefaults()
	f, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(b, &loadedConfig)
	if err != nil {
		loadedConfig = withDefaults()
		return err
	}

	return nil
}

func Get() Config {
	if !loaded {
		load()
	}

	return loadedConfig
}

func withDefaults() Config {
	return Config{
		DB: Database{
			Host: "localhost:7474",
			//Auth: false,
			//User: "",
			//Pass: "",
		},
		RootNode: "",
	}
}
