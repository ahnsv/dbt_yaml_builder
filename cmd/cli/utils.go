package main

import (
	"strings"

	"github.com/ahnsv/dbt_yaml_builder/dbt"
)

func ParseColumnStringSliceToDbtModelColumns(stringSlice []string) []dbt.DbtModelColumns {
	columns := []dbt.DbtModelColumns{}

	slice := strings.Split(stringSlice[0], ",")
	for _, columnString := range slice {
		if strings.Contains(columnString, ":") {
			splitColumnString := strings.Split(columnString, ":")
			columns = append(columns, dbt.DbtModelColumns{
				Name:        splitColumnString[0],
				Description: splitColumnString[1],
			})
			continue
		}
		columns = append(columns, dbt.DbtModelColumns{Name: columnString})
	}
	return columns
}
