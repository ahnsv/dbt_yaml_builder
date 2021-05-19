package dbt_yaml_builder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sourceYaml = ReadDbtSourceYaml("./tests/dbt_source_gitlab.yaml")
	schemaYaml = ReadDbtSchemaYaml("./tests/dbt_schema_gitlab.yaml")
)

func Test_AddNewSourceTable(t *testing.T) {
	nextSourceYaml := AddNewSourceTable(sourceYaml, "fishtown_snowplow", SourceTables{Name: "hello"})
	assert.NotNil(t, nextSourceYaml)
	var isNextTableAdded = false
	for _, source := range nextSourceYaml.Sources {
		fmt.Print(source.SourceTables)
		for _, table := range source.SourceTables {
			if table.Name == "hello" {
				isNextTableAdded = true
			}
		}
	}
	assert.True(t, isNextTableAdded)
}
