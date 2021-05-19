package dbt_yaml_builder

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func AddModelActionHandler(c *cli.Context) error {
	path := c.Args().Get(0)
	modelName := c.String("name")
	modelDescription := c.String("description")
	modelColumns := c.StringSlice("columns")
	columns := ParseColumnStringSliceToDbtModelColumns(modelColumns) // FIXME

	srcYaml := ReadDbtSchemaYaml(path)
	nextYaml := AddModel(srcYaml, Model{Name: modelName, Description: modelDescription, Columns: columns})
	out, err := yaml.Marshal(&nextYaml)
	if err != nil {
		log.Fatalf("[ModelActionHandler] Error: %v", err)
	}

	fmt.Printf("%v", string(out))
	return nil
}

func DeleteModelActionHandler(c *cli.Context) error {
	path := c.Args().Get(0)
	modelName := c.String("name")

	srcYaml := ReadDbtSchemaYaml(path)
	nextYaml := DeleteModel(srcYaml, modelName)
	out, err := yaml.Marshal(&nextYaml)
	if err != nil {
		log.Fatalf("[ModelActionHandler] Error: %v", err)
	}

	fmt.Printf("%v", string(out))
	return nil
}
