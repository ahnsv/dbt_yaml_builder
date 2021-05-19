# dbt_yaml_builder
dbt 생산성 향상을 위한 CLI 툴

```
NAME:
   dbt_yaml_builder - dbt 생산성 향상을 위한 cli 툴

USAGE:
   dbt_yaml_builder_cli [global options] command [command options] [arguments...]

COMMANDS:
   add_model, model  모델 추가
   delete_model, dm  모델 삭제
   help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Development
### Requirements
- go v1.11+
### CLI 툴 빌드
```bash
make cli-build
```
### 테스팅
```bash
make test
```

