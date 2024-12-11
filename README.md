# Go Config Loader

go-config-loader is a library that allows config to be read from yaml files. A common problem with reading from config
files in go is ensuring required values are set. Historically there are two ways of handling this:

1. Set up config variables as non-nullable and error if a field is set to its zero value.
   The issue with this is that a variable's zero value can be a valid input to an application.
2. Set up config variables as nullable and error if a field is set to null.
   This solves the issue with zero values but still means every field in the config object needs to be checked. It also
   means code that relies on the config now has to handle nullability.

Introducing go-config-loader. The library allows the setting of a required flag in the object declaration. When the flag
is set, the library will check the config yaml file for the variable and error if it isn't present. This solves the zero
values issue and means there is no need to make config variables nullable. The flag is set to false by default.

The library also has the ability to interpolate environment variable values into yaml files when the variables are
formatted in the way illustrated below. This can be useful for reading in secrets and sensitive content that should not
be stored in version control. The `api_key` variable will be set to the value of `MY_SECRET_API_KEY` when the config is
loaded.

```yaml
api_key: "$MY_SECRET_API_KEY"
```

## Getting Started

### Simple Example

```go
package main

import (
	config "github.com/loveholidays/go-config-loader"
)

type genericConfig struct {
	DummyConfig1 string `yaml:"dummy_config_1" required:"true"`
	DummyConfig2 string `yaml:"dummy_config_2"`
	DummyConfig3 string `yaml:"dummy_config_3"`
}

func main() {
	_, err := config.LoadConfiguration[genericConfig]("config.yaml")
}
```