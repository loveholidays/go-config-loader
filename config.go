package goconfigloader

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

// LoadConfiguration loads config values from the yaml file passed to it via the configPath variable. It returns a
// struct of type configT. The values of the configT struct will be set to the values inside the yaml file. If a
// value marked as required:"true" in the struct is not present in the yaml file, the function will return an error.
func LoadConfiguration[configT interface{}](configPath string) (configT, error) {
	file, err := os.ReadFile(configPath)

	var config configT
	if err != nil {
		return config, err
	}

	expandedYaml, err := expandEnvironmentVariables(file)

	if err != nil {
		return config, err
	}

	err = unmarshalAndValidate([]byte(expandedYaml), &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func expandEnvironmentVariables(file []byte) (string, error) {
	fileText := string(file)

	var missingKeys []string
	expandedYaml := os.Expand(fileText, func(key string) string {
		expanded := os.Getenv(key)
		if expanded == "" {
			missingKeys = append(missingKeys, key)
		}
		return expanded
	})

	if len(missingKeys) != 0 {
		errorMessage := "Missing required environment variables: " + strings.Join(missingKeys, ",")
		return "", errors.New(errorMessage)
	}
	return expandedYaml, nil
}

func unmarshalAndValidate(data []byte, out interface{}) error {
	var fieldsMap map[string]interface{}
	if err := yaml.Unmarshal(data, &fieldsMap); err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, out); err != nil {
		return err
	}

	return validateFields(reflect.ValueOf(out).Elem(), fieldsMap, "")
}

func validateFields(val reflect.Value, fieldsMap map[string]interface{}, prefix string) error {
	valType := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := valType.Field(i)
		yamlTag := field.Tag.Get("yaml")
		required, hasRequiredTag := field.Tag.Lookup("required")

		yamlPath := yamlTag
		if prefix != "" {
			yamlPath = prefix + "." + yamlTag
		}

		if _, found := fieldsMap[yamlTag]; !found && hasRequiredTag && required == "true" {
			return fmt.Errorf("required field '%s' is missing in YAML input", yamlPath)
		}

		if field.Type.Kind() == reflect.Struct {
			nestedFieldsMap, ok := fieldsMap[yamlTag].(map[string]interface{})
			if !ok {
				nestedFieldsMap = make(map[string]interface{}) // Handle case where the nested struct is not in the map
			}
			if err := validateFields(val.Field(i), nestedFieldsMap, yamlPath); err != nil {
				return err
			}
		}
	}
	return nil
}
