package main

import (
	"log"
	"os"

	"github.com/ahnsv/dbt_yaml_builder"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add_model",
				Aliases: []string{"m"},
				Usage:   "모델 추가",
				Action:  dbt_yaml_builder.ModelCommand,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
