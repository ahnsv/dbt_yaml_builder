cli-build:
	@cd cli && go build -o ../bin/dbt_yaml_builder_cli

test:
	go test .	