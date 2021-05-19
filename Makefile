cli-build:
	@cd cli && go build ./cmd/cli -o ../bin/dbt_yaml_builder_cli

test:
	go test ./dbt -v