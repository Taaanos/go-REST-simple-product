output "database_address" {
  value = aws_db_instance.product_database.address
}

output "database_port" {
  value = aws_db_instance.product_database.port
}

output "db_secret" {
  value = aws_ssm_parameter.db_secret.arn
}