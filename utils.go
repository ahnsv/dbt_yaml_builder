package dbt_yaml_builder

func ParseColumnStringSliceToDbtModelColumns(stringSlice []string) []DbtModelColumns {
	columns := []DbtModelColumns{}

	for _, columnString := range stringSlice {
		columns = append(columns, DbtModelColumns{Name: columnString})
	}
	return columns
}
