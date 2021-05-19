package dbt_yaml_builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sourceYaml                 = ReadDbtSourceYaml("./tests/dbt_source_gitlab.yaml")
	schemaYaml                 = ReadDbtSchemaYaml("./tests/dbt_schema_gitlab.yaml")
	expectedNewSourceTableYaml = ReadDbtSourceYaml("./tests/dbt_source_gitlab_expected.yaml")
	expectedNewSchemaModelYaml = ReadDbtSchemaYaml("./tests/dbt_schema_gitlab_expected.yaml")
	sourceName                 = "fishtown_snowplow"
	tablePayload               = SourceTable{Name: "hello"}
)

func Test_AddNewSourceTable(t *testing.T) {
	nextSourceYaml := AddNewSourceTable(sourceYaml, sourceName, tablePayload)
	assert.NotNil(t, nextSourceYaml)
	assert.Equal(t, nextSourceYaml, expectedNewSourceTableYaml)
}

func Test_DeleteSourceTable(t *testing.T) {
	deletedSourceYaml := DeleteSourceTable(expectedNewSourceTableYaml, sourceName, "hello")
	assert.Equal(t, deletedSourceYaml, sourceYaml)
}

func Test_AddModel(t *testing.T) {
	nextSchemaYaml := AddModel(schemaYaml, Model{Name: "hello"})
	assert.Equal(t, expectedNewSchemaModelYaml, nextSchemaYaml)
}
