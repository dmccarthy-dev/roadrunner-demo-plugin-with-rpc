package roadrunner_demo_plugin_with_rpc

import "fmt"

func (s *rpc) GetMessage(_ string, output *string) error {
	*output = s.plugin.cfg.Message
	s.log.Debug("foo")
	return nil
}

func (s *rpc) Add(params AddParams, output *AddResponse) error {

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
