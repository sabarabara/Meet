data "aws_region" "current" {}

resource "aws_appsync_graphql_api" "main" {
  name                = "${var.project}-${var.environment}-appsync-api"
  authentication_type = "AMAZON_COGNITO_USER_POOLS"

  user_pool_config {
    aws_region     = data.aws_region.current.name
    user_pool_id   = var.cognito_user_pool_id
    default_action = "ALLOW"
  }

  # schema = file("${path.module}/schema.graphql")
}

resource "aws_iam_role" "appsync_dynamo" {
  name = "${var.project}-${var.environment}-appsync-dynamo-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Service = "appsync.amazonaws.com"
      }
      Action = "sts:AssumeRole"
    }]
  })
}

resource "aws_iam_role_policy" "appsync_dynamo_policy" {
  role = aws_iam_role.appsync_dynamo.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action   = ["dynamodb:GetItem", "dynamodb:PutItem", "dynamodb:UpdateItem", "dynamodb:DeleteItem", "dynamodb:Query", "dynamodb:Scan"]
      Effect   = "Allow"
      Resource = [var.dynamodb_table_arn, "${var.dynamodb_table_arn}/*"]
    }]
  })
}

resource "aws_appsync_datasource" "dynamo" {
  api_id           = aws_appsync_graphql_api.main.id
  name             = "DynamoDB_DataSource"
  type             = "AMAZON_DYNAMODB"
  service_role_arn = aws_iam_role.appsync_dynamo.arn

  dynamodb_config {
    table_name = var.dynamodb_table_name
  }
}