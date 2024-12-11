package goconfigloader_test

import (
	"fmt"
	config "github.com/loveholidays/go-config-loader"
)

type configWithBoolean struct {
	DummyConfig1 bool `yaml:"dummy_config_1" required:"true"`
}

func ExampleLoadConfiguration() {
	cfg, err := config.LoadConfiguration[configWithBoolean]("config.test.bool.yaml")
	if err != nil {
		return
	}

	fmt.Printf("dummy_config_1: %v", cfg.DummyConfig1)
	//	Output: dummy_config_1: false
}
