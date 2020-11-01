package main

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/moby/buildkit/cmd/buildkitd/config"
	"github.com/pkg/errors"
)

func Load(r io.Reader) (config.Config, *toml.MetaData, error) {
	var c config.Config
	md, err := toml.DecodeReader(r, &c)
	if err != nil {
		return c, nil, errors.Wrap(err, "failed to parse config")
	}
	return c, &md, nil
}

func LoadFile(fp string) (config.Config, *toml.MetaData, error) {
	f, err := os.Open(fp)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return config.Config{}, nil, nil
		}
		return config.Config{}, nil, errors.Wrapf(err, "failed to load config from %s", fp)
	}
	defer f.Close()
	return Load(f)
}

// parseBoolOrAuto returns (nil, nil) if s is "auto"
func parseBoolOrAuto(s string) (*bool, error) {
	if s == "" || strings.ToLower(s) == "auto" {
		return nil, nil
	}
	b, err := strconv.ParseBool(s)
	return &b, err
}
