variable "project" {}
variable "environment" {}

variable "dynamodb_table_name" {}

variable "dynamodb_table_arn" {}

variable "private_subnet_ids" {
  type    = list(string)
  default = []
}

variable "alb_sg_id" {
  type    = string
  default = ""
}

variable "cognito_user_pool_id" {
  type    = string
  default = ""
}

variable "alb_dns_name" {
  type    = string
  default = ""
}


