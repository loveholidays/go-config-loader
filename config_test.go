//nolint:tagliatelle //Yaml camel case instead of snake case
package goconfigloader_test

import (
	"reflect"
	"testing"

	config "github.com/loveholidays/go-config-loader"
)

type TestConfig struct {
	InnerConfig InnerConfig `yaml:"inner"`
	SomeNumber  int         `yaml:"someNumber"`
	SomeString  string      `yaml:"someString"`
}

type InnerConfig struct {
	OtherNumber int    `yaml:"otherNumber"`
	OtherString string `yaml:"otherString"`
}

type genericConfig struct {
	DummyConfig1 string `yaml:"dummy_config_1" required:"true"`
	DummyConfig2 string `yaml:"dummy_config_2"`
	DummyConfig3 string `yaml:"dummy_config_3"`
}

type nestedConfig struct {
	DummyConfig1 nestedField `yaml:"dummy_config_1"`
	DummyConfig2 string      `yaml:"dummy_config_2"`
}

type nestedField struct {
	NestedField1 bool    `yaml:"nested_field_1" required:"true"`
	NestedField2 *string `yaml:"nested_pointer_1" required:"true"`
}

func TestExpandingEnvironmentVariablesInYAML(t *testing.T) {
	cfg, err := config.LoadConfiguration[TestConfig]("config.test.yaml")

	if cfg != nil {
		t.Error("Expected config to be nil, but it was not")
	}

	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	expectedErr := "Missing required environment variables: ENV_VAR"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestYAMLFileParsing(t *testing.T) {
	t.Setenv("ENV_VAR", "other")

	cfg, err := config.LoadConfiguration[TestConfig]("config.test.yaml")
	if err != nil {
		t.Fatal("Expected no error, but got:", err)
	}

	expected := &TestConfig{
		InnerConfig: InnerConfig{
			OtherNumber: 10,
			OtherString: "other",
		},
		SomeNumber: 12,
		SomeString: "some",
	}

	if !reflect.DeepEqual(cfg, expected) {
		t.Errorf("Expected %+v, got %+v", expected, cfg)
	}
}

func TestRequiredFields(t *testing.T) {
	t.Run("should fail to parse config when required fields are missing", func(t *testing.T) {
		t.Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")

		_, err := config.LoadConfiguration[genericConfig]("config.test.missing-field.yaml")
		if err == nil {
			t.Fatal("Expected error, but got nil")
		}

		expectedErr := "required field 'dummy_config_1' is missing in YAML input"
		if err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
		}
	})

	t.Run("should not fail when false boolean is explicitly set", func(t *testing.T) {
		t.Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")

		_, err := config.LoadConfiguration[configWithBoolean]("config.test.bool.yaml")
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}
	})

	t.Run("should not fail when false boolean is explicitly set as nested field", func(t *testing.T) {
		t.Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")

		_, err := config.LoadConfiguration[nestedConfig]("config.test.nested.yaml")
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}
	})

	t.Run("should fail when boolean is not set as nested field", func(t *testing.T) {
		t.Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")

		_, err := config.LoadConfiguration[nestedConfig]("config.test.nested.missing.yaml")
		if err == nil {
			t.Fatal("Expected error, but got nil")
		}

		expectedErr := "required field 'dummy_config_1.nested_field_1' is missing in YAML input"
		if err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
		}
	})

	t.Run("should fail when pointer is not set as nested field", func(t *testing.T) {
		t.Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")

		_, err := config.LoadConfiguration[nestedConfig]("config.test.nested.missing-pointer.yaml")
		if err == nil {
			t.Fatal("Expected error, but got nil")
		}

		expectedErr := "required field 'dummy_config_1.nested_pointer_1' is missing in YAML input"
		if err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%s'", expectedErr, err.Error())
		}
	})
}
