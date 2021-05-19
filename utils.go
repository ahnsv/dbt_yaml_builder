package dbt_yaml_builder

import (
	"strings"
)

func ParseColumnStringSliceToDbtModelColumns(stringSlice []string) []DbtModelColumns {
	columns := []DbtModelColumns{}

	slice := strings.Split(stringSlice[0], ",")
	for _, columnString := range slice {
		if strings.Contains(columnString, ":") {
			splitColumnString := strings.Split(columnString, ":")
			columns = append(columns, DbtModelColumns{
				Name:        splitColumnString[0],
				Description: splitColumnString[1],
			})
			continue
		}
		columns = append(columns, DbtModelColumns{Name: columnString})
	}
	return columns
}
