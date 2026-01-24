## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_api"></a> [api](#module\_api) | ../modules/api | n/a |
| <a name="module_compute"></a> [compute](#module\_compute) | ../modules/compute | n/a |
| <a name="module_database"></a> [database](#module\_database) | ../modules/database | n/a |
| <a name="module_frontend"></a> [frontend](#module\_frontend) | ../modules/frontend | n/a |
| <a name="module_network"></a> [network](#module\_network) | ../modules/network | n/a |
| <a name="module_security"></a> [security](#module\_security) | ../modules/security | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_environment"></a> [environment](#input\_environment) | n/a | `string` | n/a | yes |
| <a name="input_postgres_database"></a> [postgres\_database](#input\_postgres\_database) | n/a | `string` | n/a | yes |
| <a name="input_postgres_host"></a> [postgres\_host](#input\_postgres\_host) | n/a | `string` | n/a | yes |
| <a name="input_postgres_password"></a> [postgres\_password](#input\_postgres\_password) | n/a | `string` | n/a | yes |
| <a name="input_postgres_username"></a> [postgres\_username](#input\_postgres\_username) | n/a | `string` | n/a | yes |
| <a name="input_project"></a> [project](#input\_project) | n/a | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_dynamodb_table_arn"></a> [dynamodb\_table\_arn](#output\_dynamodb\_table\_arn) | n/a |
| <a name="output_postgres_address"></a> [postgres\_address](#output\_postgres\_address) | n/a |
