output "postgres_address" {
  value = module.database.postgres_address
}

output "dynamodb_table_arn" {
  value = module.database.dynamodb_table_arn
}