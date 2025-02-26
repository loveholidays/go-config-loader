//nolint:tagliatelle //Yaml camel case instead of snake case
package goconfigloader_test

import (
	"fmt"
	config "github.com/loveholidays/go-config-loader"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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

func TestLoadConfiguration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shared Config")
}

var _ = Describe("Shared Config", func() {

	Describe("Expanding environment variables in YAML", func() {
		It("returns error for missing variables", func() {
			_, err := config.LoadConfiguration[TestConfig]("config.test.yaml")

			Expect(err.Error()).To(Equal("Missing required environment variables: ENV_VAR"))
		})
	})

	Describe("YAML file parsing", func() {
		It("unmarshalls into cfg", func() {
			os.Clearenv()
			_ = os.Setenv("ENV_VAR", "other")
			cfg, err := config.LoadConfiguration[TestConfig]("config.test.yaml")
			Expect(err).To(BeNil())

			expected := TestConfig{
				InnerConfig: InnerConfig{
					OtherNumber: 10,
					OtherString: "other",
				},
				SomeNumber: 12,
				SomeString: "some",
			}

			Expect(cfg).To(Equal(expected))
		})

	})

	Describe("Required fields", func() {
		It("should fail to parse config when required fields are missing", func() {
			os.Clearenv()
			GinkgoT().Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")
			_, err := config.LoadConfiguration[genericConfig]("config.test.missing-field.yaml")

			Expect(err).To(HaveOccurred())
			fmt.Printf("%s", err.Error())
			Expect(err.Error()).To(Equal("required field 'dummy_config_1' is missing in YAML input"))
		})

		It("should not fail when false boolean is explicitly set", func() {
			os.Clearenv()
			GinkgoT().Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")
			_, err := config.LoadConfiguration[configWithBoolean]("config.test.bool.yaml")

			Expect(err).ToNot(HaveOccurred())
		})

		It("should not fail when false boolean is explicitly set as nested field", func() {
			os.Clearenv()
			GinkgoT().Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")
			_, err := config.LoadConfiguration[nestedConfig]("config.test.nested.yaml")

			Expect(err).ToNot(HaveOccurred())
		})

		It("should fail when boolean is not set as nested field", func() {
			os.Clearenv()
			GinkgoT().Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")
			_, err := config.LoadConfiguration[nestedConfig]("config.test.nested.missing.yaml")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("required field 'dummy_config_1.nested_field_1' is missing in YAML input"))
		})

		It("should fail when pointer is not set as nested field", func() {
			os.Clearenv()
			GinkgoT().Setenv("DUMMY_CONFIG_2_COMES_FROM_ENV", "set from env")
			_, err := config.LoadConfiguration[nestedConfig]("config.test.nested.missing-pointer.yaml")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("required field 'dummy_config_1.nested_pointer_1' is missing in YAML input"))
		})
	})
})
