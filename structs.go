package dbt_yaml_builder

// Tables
type SourceTable struct {
	Name        string               `yaml:"name"`
	Identifier  string               `yaml:"identifier"`
	Description string               `yaml:"description"`
	Columns     []SourceTableColumns `yaml:"columns"`
	Freshness   interface{}          `yaml:"freshness"`
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
type SourceTableColumns struct {
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
	Sources []Source `yaml:"sources"`
	Version int      `yaml:"version"`
}

// Source
type Source struct {
	SourceTables  []SourceTable `yaml:"tables"`
	Name          string        `yaml:"name"`
	Description   string        `yaml:"description"`
	Database      string        `yaml:"database"`
	Schema        string        `yaml:"schema"`
	Loader        string        `yaml:"loader"`
	LoadedAtField string        `yaml:"loaded_at_field"`
	Quoting       SourceQuoting `yaml:"quoting"`
}

// DbtSchemaYaml
type DbtSchemaYaml struct {
	Version int      `yaml:"version"`
	Models  []Models `yaml:"models"`
}

// Models
type Models struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Columns     []DbtModelColumns `yaml:"columns"`
}

// DbtModelColumns
type DbtModelColumns struct {
	Description string   `yaml:"description"`
	Name        string   `yaml:"name"`
	Tests       []string `yaml:"tests"`
}
