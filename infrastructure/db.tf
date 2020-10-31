resource "aws_security_group" "db_allow_all" {
  name        = "db_allow_all"
  description = "Allow all inbound traffic"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description = "Postgres from all"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Terraform = true
    Name      = "db_allow_all"
  }
}

resource "aws_db_instance" "product_database" {
  allocated_storage      = 20
  storage_type           = "gp2"
  engine                 = "postgres"
  engine_version         = "12.3"
  instance_class         = "db.t2.micro"
  name                   = "productDB"
  username               = var.db_username
  password               = var.db_password
  publicly_accessible    = true
  vpc_security_group_ids = [aws_security_group.db_allow_all.id]
}


output "database_address" {
  value = aws_db_instance.product_database.address
}

output "database_port" {
  value = aws_db_instance.product_database.port
}
