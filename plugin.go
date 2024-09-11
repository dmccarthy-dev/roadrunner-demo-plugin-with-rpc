package roadrunner_demo_plugin_with_rpc

import (
	"context"
	"github.com/roadrunner-server/errors"
	"go.uber.org/zap"
)

const PluginName = "plugin_with_rpc"

type Plugin struct {
	cfg    *Config
	logger *zap.Logger
}

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshal it into a Struct.
	UnmarshalKey(name string, out any) error
	// Has checks if config section exists.
	Has(name string) bool
}

type Logger interface { // <-- logger plugin implements
	NamedLogger(name string) *zap.Logger
}

func (p *Plugin) Init(cfg Configurer, log Logger) error {
	p.logger = log.NamedLogger("roadrunner_demo_plugin_with_rpc")
	p.logger.Info("start initializing demo plugin")
	const op = errors.Op("custom_plugin_init")

	if !cfg.Has(PluginName) {
		return errors.E(op, errors.Disabled)
	}

	err := cfg.UnmarshalKey(PluginName, &p.cfg)
	if err != nil {
		// Error will stop execution
		return errors.E(op, err)
	}

	p.cfg.InitDefaults()

	p.logger.Info("end initializing demo plugin")
	return nil
}

func (p *Plugin) Serve() chan error {
	return nil
}

func (p *Plugin) Stop(_ context.Context) error {
	p.logger.Info("Stopping plugin")
	return nil
}

func (p *Plugin) Name() string {
	return PluginName
}
