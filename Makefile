cli-build:
	@go build -o ../bin/dbt_yaml_builder_cli ./cmd/cli 

test:
	go test ./dbt -v