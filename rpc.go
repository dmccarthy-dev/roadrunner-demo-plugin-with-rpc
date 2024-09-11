package roadrunner_demo_plugin_with_rpc

import (
	"fmt"
	"go.uber.org/zap"
)

type rpc struct {
	plugin *Plugin
	log    *zap.Logger
}

func (p *Plugin) RPC() any {
	return &rpc{plugin: p, log: p.logger}
}

func (r *rpc) GetMessage(_ string, output *string) error {
	*output = r.plugin.cfg.Message
	r.log.Debug("foo")
	return nil
}

func (r *rpc) Add(params AddParams, output *AddResponse) error {

	fmt.Println("params.Num1", params.Num1)
	fmt.Println("params.Num2", params.Num2)

	res := params.Num1 + params.Num2

	*output = AddResponse{Result: res}
	return nil
}

type AddParams struct {
	Num1 int `mapstructure:"num1"`
	Num2 int `mapstructure:"num2"`
}

type AddResponse struct {
	Result int `json:"result"`
}
