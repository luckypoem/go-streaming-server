package conf

import "github.com/BurntSushi/toml"

type Config struct {
	Host          string `toml:"HOST"`
	Port          int    `toml:"PORT"`
	MaxConnection int    `toml:"MAX_CONNECTION"`
	VideoDir      string `toml:"VIDEO_DIR"`
	MaxUploadSize int64  `toml:"MAX_UPLOADSIZE"`
}

func LoadConfigFromFile(filepath string) (*Config, error) {
	var conf Config

	_, err := toml.DecodeFile(filepath, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
