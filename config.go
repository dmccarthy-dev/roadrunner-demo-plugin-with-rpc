package roadrunner_demo_plugin_with_rpc

type Config struct {
	Message string `mapstructure:"message"`
}

// InitDefaults .. You can also initialize some defaults values for config keys
func (cfg *Config) InitDefaults() {
	if cfg.Message == "" {
		cfg.Message = "default hello world"
	}
}
