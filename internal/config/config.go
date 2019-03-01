package config

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/tserkov/google-ddns-client/internal/record"
)

type Config struct {
	Records    []record.Record `json:"records" yaml:"records"`
	IPProvider string          `json:"ip_provider" yaml:"ip_provider"`
	Verbose    bool            `json:"verbose" yaml:"verbose"`
}

const (
	defaultIPProvider string = "icanhazip.com"
)

func Parse() (*Config, error) {
	filename := flag.String("c", "", "location of config file")
	verbose := flag.Bool("v", false, "be verbose")
	flag.Parse()

	if filename == nil || *filename == "" {
		return nil, errors.New("must specify config file location within -c")
	}

	c, err := parseFile(*filename)
	if err != nil {
		return nil, err
	}

	for _, r := range c.Records {
		if err := r.Valid(); err != nil {
			return nil, err
		}
	}

	c.Verbose = *verbose

	return c, nil
}

func parseFile(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(filename)
	switch ext {
	case ".json":
		return parseJSON(data)
	case ".yaml":
		return parseYAML(data)
	case ".yml":
		return parseYAML(data)
	}

	return nil, errors.New("config has unsupported format: " + ext)
}

func parseJSON(data []byte) (*Config, error) {
	c := Config{}

	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

func parseYAML(data []byte) (*Config, error) {
	c := Config{}

	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
