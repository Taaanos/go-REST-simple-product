resource "aws_ecs_service" "product_service" {
  name            = "product-service"
  cluster         = aws_ecs_cluster.window_cluster.id
  task_definition = aws_ecs_task_definition.product_service.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  network_configuration {
    subnets          = data.aws_subnet_ids.default.ids
    security_groups  = [aws_security_group.allow_http_product_service.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.product_service.arn
    container_name   = "product-service-container"
    container_port   = 3010
  }

  tags = {
    Terraform = true
  }
}

resource "aws_security_group" "allow_http_product_service" {
  name        = "allow_http_product_service"
  description = "Allow http indound traffic at port 3010"

  ingress {
    description = "TCP from all"
    from_port   = 3010
    to_port     = 3010
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Terraform = true
    Name      = "allow_http_product_service"
  }
}