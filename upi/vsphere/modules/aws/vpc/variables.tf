variable "machine_cidr" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "availability_zones" {
  type        = list(string)
  description = "The availability zones in which to provision subnets."
}

variable "cidr_block" {
  type = string
}

variable "cluster_id" {
  type = string
}

variable "private_control_plane_endpoints" {
  description = "If set to true, private-facing ingress resources are created."
  default     = true
}

variable "public_control_plane_endpoints" {
  description = "If set to true, public-facing ingress resources are created."
  default     = true
}

variable "region" {
  type        = string
  description = "The target AWS region for the cluster."
}

variable "tags" {
  type        = map(string)
  default     = {}
  description = "AWS tags to be applied to created resources."
}

variable "public_subnet_id" {
  type = list(string)
}

variable "private_subnet_id" {
  type = list(string)
}

