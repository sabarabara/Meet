resource "aws_secretsmanager_secret" "postgres" {
  name                    = "${var.project}-${var.environment}-postgres-credentials"
  description             = "PostgreSQL credentials for ${var.project} in ${var.environment} environment"
  recovery_window_in_days = 0

  tags = {
    Name = "${var.project}-${var.environment}-postgres-credentials"
  }
}

resource "aws_secretsmanager_secret_version" "postgres" {
  secret_id = aws_secretsmanager_secret.postgres.id
  secret_string = jsonencode({
    username = var.postgres_username
    password = var.postgres_password
    host     = var.postgres_host
    database = var.postgres_database
  })

  lifecycle {
    ignore_changes = [
      secret_string
    ]
  }
}