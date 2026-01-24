resource "aws_apigatewayv2_api" "http_gw" {
  name          = "my-app-http-api"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "default" {
  api_id      = aws_apigatewayv2_api.http_gw.id
  name        = "$default"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_logs.arn
    format          = "$context.requestId $context.ip $context.requestTime $context.httpMethod $context.routeKey $context.status $context.responseLength $context.integrationLatency"
  }
}

resource "aws_cloudwatch_log_group" "api_logs" {
  name              = "/aws/apigateway/${var.project}-${var.environment}"
  retention_in_days = 7

  tags = {
    Name = "${var.project}-${var.environment}-api-logs"
  }
}

resource "aws_apigatewayv2_vpc_link" "main" {
  name               = "${var.project}-${var.environment}-vpc-link"
  subnet_ids         = var.private_subnet_ids
  security_group_ids = var.alb_sg_id != "" ? [var.alb_sg_id] : []
}

resource "aws_apigatewayv2_integration" "alb" {
  api_id                 = aws_apigatewayv2_api.http_gw.id
  integration_type       = "HTTP_PROXY"
  integration_uri        = var.alb_dns_name
  integration_method     = "ANY"
  payload_format_version = "1.0"
  connection_type        = "VPC_LINK"
  connection_id          = aws_apigatewayv2_vpc_link.main.id
}

resource "aws_apigatewayv2_route" "default" {
  api_id    = aws_apigatewayv2_api.http_gw.id
  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.alb.id}"
}