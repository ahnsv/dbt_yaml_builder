package main

import (
	"fmt"
	"log"

	"github.com/ahnsv/dbt_yaml_builder/dbt"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func AddModelActionHandler(c *cli.Context) error {
	path := c.Args().Get(0)
	modelName := c.String("name")
	modelDescription := c.String("description")
	modelColumns := c.StringSlice("columns")
	columns := ParseColumnStringSliceToDbtModelColumns(modelColumns) // FIXME

	srcYaml := dbt.ReadDbtSchemaYaml(path)
	nextYaml := dbt.AddModel(srcYaml, dbt.Model{Name: modelName, Description: modelDescription, Columns: columns})
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

	srcYaml := dbt.ReadDbtSchemaYaml(path)
	nextYaml := dbt.DeleteModel(srcYaml, modelName)
	out, err := yaml.Marshal(&nextYaml)
	if err != nil {
		log.Fatalf("[ModelActionHandler] Error: %v", err)
	}

	fmt.Printf("%v", string(out))
	return nil
}
