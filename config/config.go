package config

// Configuration models env configuration
type Configuration struct {
	BasePath          string `envconfig:"base_path" default:"/kuraifu"`
	BasePort          string `envconfig:"base_port" default:"8080"`
	LineChannelSecret string `envconfig:"line_channel_secret" required:"true"`
	LineAccessToken   string `envconfig:"line_access_token" required:"true"`
}
