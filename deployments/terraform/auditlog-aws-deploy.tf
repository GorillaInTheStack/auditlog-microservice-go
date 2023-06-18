# AWS
provider "aws" {
  access_key = var.access_key
  secret_key = var.secret_key
  region     = var.region
}

# VPC resource
resource "aws_vpc" "auditlog_microservice" {
  cidr_block = var.cidr_block
}

# Subnet resource
resource "aws_subnet" "auditlog_microservice" {
  vpc_id                  = aws_vpc.auditlog_microservice.id
  cidr_block              = var.subnet_cidr_block
  availability_zone       = var.availability_zone
}

# Security group resource
resource "aws_security_group" "auditlog_microservice" {
  name        = "auditlog-microservice-security-group"
  description = "Security group for the auditlog microservice"

  vpc_id = aws_vpc.auditlog_microservice.id

  # Inbound rule for HTTP traffic
  ingress {
    from_port   = 0
    to_port     = 6969
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Inbound rule for SSH traffic
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.ssh_cidr_block]
  }

  # Outbound rule to allow all traffic
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# EC2 instance resource
resource "aws_instance" "auditlog_microservice" {
  ami           = var.ami
  instance_type = var.instance_type
  subnet_id     = aws_subnet.auditlog_microservice.id

  vpc_security_group_ids = [aws_security_group.auditlog_microservice.id]

  user_data = var.user_data_script

  tags = {
    Name = "auditlog-microservice-instance"
  }
}
