resource "aws_security_group" "app" {
  name        = "${var.project}-${var.environment}-app-sg"
  vpc_id      = var.vpc_id
  description = "Security group for application servers"

  tags = {
    Name = "${var.project}-${var.environment}-app-sg"
  }
}

resource "aws_security_group_rule" "app_in_http" {
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  security_group_id = aws_security_group.app.id
  cidr_blocks       = []
  description       = "Allow HTTP from ALB"
}

resource "aws_security_group_rule" "app_in_https" {
  type              = "ingress"
  from_port         = 443
  to_port           = 443
  protocol          = "tcp"
  security_group_id = aws_security_group.app.id
  cidr_blocks       = []
  description       = "Allow HTTPS from ALB"
}

#tfsec:ignore=aws-ec2-no-public-egress-sgr
resource "aws_security_group_rule" "app_out_https" {
  type              = "egress"
  from_port         = 443
  to_port           = 443
  protocol          = "tcp"
  security_group_id = aws_security_group.app.id
  cidr_blocks       = ["0.0.0.0/0"] # Required for external APIs (npm, pip, maven)
  description       = "Allow HTTPS outbound for package downloads"
}

resource "aws_security_group" "database" {
  name        = "${var.project}-${var.environment}-db-sg"
  vpc_id      = var.vpc_id
  description = "Security group for database servers"

  tags = {
    Name = "${var.project}-${var.environment}-db-sg"
  }
}

resource "aws_security_group_rule" "db_in_tcp_5432" {
  type                     = "ingress"
  from_port                = 5432
  to_port                  = 5432
  protocol                 = "tcp"
  security_group_id        = aws_security_group.database.id
  source_security_group_id = aws_security_group.app.id
  description              = "Allow PostgreSQL from application security group"
}