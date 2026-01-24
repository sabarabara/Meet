resource "aws_db_instance" "postgres" {
  identifier                            = "${var.project}-${var.environment}-postgres-db"
  engine                                = "postgres"
  instance_class                        = "db.t3.micro"
  allocated_storage                     = 20
  storage_encrypted                     = true
  db_name                               = var.postgres_database
  username                              = var.postgres_username
  password                              = var.postgres_password
  vpc_security_group_ids                = var.db_security_group_ids
  db_subnet_group_name                  = aws_db_subnet_group.postgres.name
  multi_az                              = false
  publicly_accessible                   = false
  skip_final_snapshot                   = true
  backup_retention_period               = 7
  iam_database_authentication_enabled   = true
  deletion_protection                   = true
  performance_insights_retention_period = 7

  tags = {
    Name = "${var.project}-${var.environment}-postgres-db"
  }
}

resource "aws_db_subnet_group" "postgres" {
  name       = "${var.project}-${var.environment}-rds-subnet-group"
  subnet_ids = var.db_subnet_ids

  tags = {
    Name = "${var.project}-${var.environment}-rds-subnet-group"
  }
}

resource "aws_db_parameter_group" "postgres" {
  name        = "${var.project}-${var.environment}-rds-parameter-group"
  family      = "postgres12"
  description = "Custom parameter group for PostgreSQL"

  parameter {
    name  = "max_connections"
    value = "100"
  }
  parameter {
    name  = "character_set_server"
    value = "utf8mb4"
  }

  parameter {
    name  = "collation_server"
    value = "utf8mb4_unicode_ci"
  }

  parameter {
    name         = "lower_case_table_names"
    value        = "0"
    apply_method = "pending-reboot"
  }

  parameter {
    name  = "require_secure_transport"
    value = "0"
  }

  parameter {
    name  = "time_zone"
    value = "Asia/Tokyo"
  }

  tags = {
    Name = "${var.project}-${var.environment}-rds-parameter-group"
  }
}