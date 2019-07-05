package conf

import "github.com/BurntSushi/toml"

// 定义配置结构体
type Config struct {
	Host string `toml:"HOST"`
	Port int    `toml:"PORT"`
}

// 加载配置文件
func LoadConfigFromFile(filePath string) (*Config, error) {
	var config Config
	// 使用toml解析器解析toml文件
	_, err := toml.DecodeFile(filePath, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
