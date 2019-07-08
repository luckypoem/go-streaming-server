package conf

type Config struct {
	Host          string `toml:"HOST"`
	Port          int    `toml:"PORT"`
	MaxConnection int    `toml:"MAX_CONNECTION"`
	VideoDir      string `toml:"VIDEO_DIR"`
	MaxUploadSize int    `toml"MAX_UPLOAD_SIZE"`
}
