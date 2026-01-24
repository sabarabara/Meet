terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region                   = "ap-northeast-1"
  shared_credentials_files = ["~/.aws/credentials"]
  profile                  = "dev"
}

module "network" {
  source      = "../modules/network"
  project     = var.project
  environment = var.environment
}

module "security" {
  source            = "../modules/security"
  project           = var.project
  environment       = var.environment
  postgres_username = var.postgres_username
  postgres_password = var.postgres_password
  postgres_host     = var.postgres_host
  postgres_database = var.postgres_database
  vpc_id            = module.network.vpc_id
}

module "compute" {
  source      = "../modules/compute"
  project     = var.project
  environment = var.environment
}

module "frontend" {
  source      = "../modules/frontend"
  project     = var.project
  environment = var.environment
}

module "database" {
  source            = "../modules/database"
  project           = var.project
  environment       = var.environment
  postgres_username = var.postgres_username
  postgres_password = var.postgres_password
  postgres_host     = var.postgres_host
  postgres_database = var.postgres_database
}

module "api" {
  source              = "../modules/api"
  project             = var.project
  environment         = var.environment
  dynamodb_table_name = module.database.dynamodb_table_name
  dynamodb_table_arn  = module.database.dynamodb_table_arn
}
