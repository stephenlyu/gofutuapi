package futuapi

import (
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

type Config struct {
	PrivateKeyFile string    `json:"private_key_file"`
	ClientInfo     string    `json:"client_info"`
	IP             string            `json:"ip"`
	Port           uint16            `json:"port"`
	UserId 		   uint64 	`json:"user_id"`

	PrivateKey     string    `json:"-"`
}

func (config *Config) Load(configFile string) error {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		return err
	}
	var privateKeyFile = config.PrivateKeyFile
	if !filepath.IsAbs(privateKeyFile) {
		privateKeyFile = filepath.Join(filepath.Dir(configFile), privateKeyFile)
	}
	bytes, err = ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return err
	}
	config.PrivateKey = string(bytes)
	return nil
}
