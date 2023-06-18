variable "access_key" {
  description = "AWS access key"
  type        = string
  default = "value"
}

variable "secret_key" {
  description = "AWS secret key"
  type        = string
  default = "value"
}

variable "region" {
  description = "AWS region where the resources will be provisioned"
  type        = string
  default     = "eu-north-1"
}

variable "cidr_block" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "subnet_cidr_block" {
  description = "CIDR block for subnet"
  type        = string
  default     = "10.0.1.0/24"
}

variable "availability_zone" {
  description = "Availability zone for subnet"
  type        = string
  default     = "eu-north-1a"
}

variable "ssh_cidr_block" {
  description = "CIDR block for SSH access"
  type        = string
  default = "0.0.0.0/32"
}

variable "ami" {
  description = "AMI ID for the EC2 instance"
  type        = string
  default     = "ami-0c55b159cbfafe1f0"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t2.micro"
}

variable "user_data_script" {
  description = "Configuration management for the auditlog microservice"
  type        = string
  default     = <<-EOF
    #!/bin/bash
    # Install dependencies and start the auditlog microservice
    apt-get update
    apt-get install -y git golang
    git clone https://gitlab.com/Sam66ish/auditlog-microservice # Will not work currently because repo is private.
    cd auditlog-microservice
    go run main.go
  EOF
}
