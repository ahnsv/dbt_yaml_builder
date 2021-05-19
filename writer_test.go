package dbt_yaml_builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sourceYaml                 = ReadDbtSourceYaml("./tests/dbt_source_gitlab.yaml")
	schemaYaml                 = ReadDbtSchemaYaml("./tests/dbt_schema_gitlab.yaml")
	expectedNewSourceTableYaml = ReadDbtSourceYaml("./tests/dbt_source_gitlab_expected.yaml")
)

func Test_AddNewSourceTable(t *testing.T) {
	nextSourceYaml := AddNewSourceTable(sourceYaml, "fishtown_snowplow", SourceTable{Name: "hello"})
	assert.NotNil(t, nextSourceYaml)
	assert.Equal(t, nextSourceYaml, expectedNewSourceTableYaml)
}
