package dbt_yaml_builder

// Tables
type SourceTables struct {
	Name              string              `yaml:"name"`
	Identifier        string              `yaml:"identifier"`
	Description       string              `yaml:"description"`
	SourceTableColumn []SourceTableColumn `yaml:"columns"`
	Freshness         interface{}         `yaml:"freshness"`
}

type Freshness struct {
	ErrorAfter ErrorAfter `yaml:"error_after"`
	WarnAfter  WarnAfter  `yaml:"warn_after"`
}

// ErrorAfter
type ErrorAfter struct {
	Count  int    `yaml:"count"`
	Period string `yaml:"period"`
}

// WarnAfter
type WarnAfter struct {
	Count  int    `yaml:"count"`
	Period string `yaml:"period"`
}

// Columns
type SourceTableColumn struct {
	Name  string   `yaml:"name"`
	Tests []string `yaml:"tests"`
}

// Quoting
type SourceQuoting struct {
	Database   bool `yaml:"database"`
	Schema     bool `yaml:"schema"`
	Identifier bool `yaml:"identifier"`
}

// Yaml2Go
type DbtSourceYaml struct {
	Sources []Sources `yaml:"sources"`
	Version int       `yaml:"version"`
}

// Sources
type Sources struct {
	SourceTables  []SourceTables `yaml:"tables"`
	Name          string         `yaml:"name"`
	Description   string         `yaml:"description"`
	Database      string         `yaml:"database"`
	Schema        string         `yaml:"schema"`
	Loader        string         `yaml:"loader"`
	LoadedAtField string         `yaml:"loaded_at_field"`
	SourceQuoting SourceQuoting  `yaml:"quoting"`
}
