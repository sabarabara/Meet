resource "aws_iam_policy" "secret_manager" {
  name        = "${var.project}-${var.environment}-ecs-secret-manager-read"
  description = "Policy to allow access to AWS Secrets Manager"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "secretsmanager:GetSecretValue",
          "kms:Decrypt"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
  tags = {
    Name = "${var.project}-${var.environment}-ecs-secret-manager-read-policy"
  }
}

resource "aws_iam_role" "ecs_task_execution" {
  name = "${var.project}-${var.environment}-ecs-task-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name = "${var.project}-${var.environment}-ecs-task-execution-role"
  }
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_attachment" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = aws_iam_policy.secret_manager.arn
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_amazon_ecs" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_cloudwatch_log_group" "app" {
  name              = "/ecs/${var.project}-${var.environment}-app"
  retention_in_days = 7

  tags = {
    Name = "${var.project}-${var.environment}-app-log-group"
  }
}
data "aws_region" "current" {}
resource "aws_ecs_cluster" "main" {
  name = "${var.project}-${var.environment}-ecs-cluster"
  setting {
    name  = "containerInsights"
    value = "enabled"
  }
  tags = {
    Name = "${var.project}-${var.environment}-ecs-cluster"
  }
}

resource "aws_ecs_task_definition" "app" {
  family                   = "${var.project}-${var.environment}-app-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_task_execution.arn

  container_definitions = jsonencode([
    {
      name  = "webapp"
      image = "${aws_ecr_repository.webapp.repository_url}:latest"
      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
          protocol      = "tcp"
        }
      ]

      Secrets = [
        {
          name      = "POSTGRES_HOST"
          valueFrom = "${var.postgres_secret_arn}:hostname::"
        },
        {
          name      = "POSTGRES_USER"
          valueFrom = "${var.postgres_secret_arn}:username::"
        },
        {
          name      = "POSTGRES_PASSWORD"
          valueFrom = "${var.postgres_secret_arn}:password::"
        },
        {
          name      = "POSTGRES_DATABASE"
          valueFrom = "${var.postgres_secret_arn}:database::"
        },
      ]

      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = aws_cloudwatch_log_group.app.name
          "awslogs-region"        = data.aws_region.current.name
          "awslogs-stream-prefix" = "ecs"
        }
      }
    }
  ])

  tags = {
    Name = "${var.project}-${var.environment}-app-task-definition"
  }
}

resource "aws_ecs_service" "app" {
  name            = "${var.project}-${var.environment}-app-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.app.arn
  launch_type     = "FARGATE"
  desired_count   = 0

  network_configuration {
    subnets          = var.app_subnet_ids
    security_groups  = var.ecs_security_group_id != "" ? [var.ecs_security_group_id] : []
    assign_public_ip = false
  }

  depends_on = [
    aws_iam_role_policy_attachment.ecs_task_execution_attachment,
    aws_iam_role_policy_attachment.ecs_task_execution_amazon_ecs
  ]

  load_balancer {
    target_group_arn = var.app_target_group_arn
    container_name   = "webapp"
    container_port   = 80
  }

  tags = {
    Name = "${var.project}-${var.environment}-app-service"
  }
}