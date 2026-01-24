variable "project" {}

variable "environment" {}

variable "postgres_database" {}

variable "postgres_username" {}

variable "postgres_password" {}

variable "postgres_host" {}

variable "db_subnet_ids" {
  type    = list(string)
  default = []
}

variable "db_security_group_ids" {
  type    = list(string)
  default = []
}
