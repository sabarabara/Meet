variable "project" {}
variable "environment" {}
variable "postgres_secret_arn" {
  type    = string
  default = ""
}
variable "public_subnet_ids" {
  type    = list(string)
  default = []
}
variable "vpc_id" {
  type    = string
  default = ""
}
variable "alb_sg_id" {
  type    = string
  default = ""
}
variable "app_subnet_ids" {
  type    = list(string)
  default = []
}
variable "ecs_security_group_id" {
  type    = string
  default = ""
}
variable "app_target_group_arn" {
  type    = string
  default = ""
}



