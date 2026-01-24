data "aws_region" "current" {}

resource "aws_vpc_endpoint" "s3" {
  vpc_id            = aws_vpc.main.id
  service_name      = "com.amazonaws.${data.aws_region.current.name}.s3"
  vpc_endpoint_type = "Gateway"

  route_table_ids = [
    aws_route_table.public.id,
  ]

  tags = {
    Name = "${var.project}-${var.environment}-s3-endpoint"
  }

}

resource "aws_vpc_endpoint" "ecr_dck" {
  vpc_id             = aws_vpc.main.id
  service_name       = "com.amazonaws.${data.aws_region.current.name}.ecr.dkr"
  vpc_endpoint_type  = "Interface"
  subnet_ids         = aws_subnet.app[*].id
  security_group_ids = []

  tags = {
    Name = "${var.project}-${var.environment}-ecr-dck-endpoint"
  }
}

resource "aws_vpc_endpoint" "ecr_api" {
  vpc_id             = aws_vpc.main.id
  service_name       = "com.amazonaws.${data.aws_region.current.name}.ecr.api"
  vpc_endpoint_type  = "Interface"
  subnet_ids         = aws_subnet.app[*].id
  security_group_ids = []

  tags = {
    Name = "${var.project}-${var.environment}-ecr-api-endpoint"
  }
}

resource "aws_vpc_endpoint" "cloudwatch" {
  vpc_id             = aws_vpc.main.id
  service_name       = "com.amazonaws.${data.aws_region.current.name}.logs"
  vpc_endpoint_type  = "Interface"
  subnet_ids         = aws_subnet.app[*].id
  security_group_ids = []

  tags = {
    Name = "${var.project}-${var.environment}-cloudwatch-endpoint"
  }
}

resource "aws_vpc_endpoint" "ssm" {
  vpc_id             = aws_vpc.main.id
  service_name       = "com.amazonaws.${data.aws_region.current.name}.ssm"
  vpc_endpoint_type  = "Interface"
  subnet_ids         = aws_subnet.app[*].id
  security_group_ids = []

  tags = {
    Name = "${var.project}-${var.environment}-ssm-endpoint"
  }
}

resource "aws_vpc_endpoint" "secret_manager" {
  vpc_id             = aws_vpc.main.id
  service_name       = "com.amazonaws.${data.aws_region.current.name}.secretsmanager"
  vpc_endpoint_type  = "Interface"
  subnet_ids         = aws_subnet.app[*].id
  security_group_ids = []

  tags = {
    Name = "${var.project}-${var.environment}-secrets-manager-endpoint"
  }
}

resource "aws_vpc_endpoint" "ssm_manager" {
  vpc_id             = aws_vpc.main.id
  service_name       = "com.amazonaws.${data.aws_region.current.name}.ssmmessages"
  vpc_endpoint_type  = "Interface"
  subnet_ids         = aws_subnet.app[*].id
  security_group_ids = []

  tags = {
    Name = "${var.project}-${var.environment}-ssm-messages-endpoint"
  }
}