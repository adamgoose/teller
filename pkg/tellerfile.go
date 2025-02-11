package pkg

import (
	"io/ioutil"

	"github.com/spectralops/teller/pkg/core"
	"gopkg.in/yaml.v2"
)

type TellerFile struct {
	Opts       map[string]string        `yaml:"opts,omitempty"`
	Confirm    string                   `yaml:"confirm,omitempty"`
	Project    string                   `yaml:"project,omitempty"`
	CarryEnv   bool                     `yaml:"carry_env,omitempty"`
	Providers  map[string]MappingConfig `yaml:"providers,omitempty"`
	LoadedFrom string
}

type MappingConfig struct {
	EnvMapping *core.KeyPath            `yaml:"env_sync,omitempty"`
	Env        *map[string]core.KeyPath `yaml:"env,omitempty"`
}

func NewTellerFile(s string) (*TellerFile, error) {
	yamlFile, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}
	t := &TellerFile{}
	err = yaml.Unmarshal(yamlFile, t)
	if err != nil {
		return nil, err
	}
	t.LoadedFrom = s
	return t, nil
}
