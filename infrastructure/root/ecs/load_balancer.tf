resource "aws_lb" "window_lb" {
  name               = "window-lb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.http_allow_all.id]
  subnets            = data.aws_subnet_ids.default.ids

  enable_deletion_protection = false

  tags = {
    Environment = "production"
    Terraform   = true
  }
}

resource "aws_lb_listener" "window_lb_http" {
  load_balancer_arn = aws_lb.window_lb.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "application/json"
      message_body = "{ \"load_balancer\": \"Oops!\"}"
      status_code  = "200"
    }
  }
}

resource "aws_lb_listener_rule" "product_service" {
  listener_arn = aws_lb_listener.window_lb_http.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.product_service.arn
  }

  condition {
    path_pattern {
      values = ["/"]
    }
  }
}

resource "aws_security_group" "http_allow_all" {
  name        = "http_allow_all"
  description = "Allow all inbound traffic"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description = "TCP from all"
    from_port   = 80
    to_port     = 80
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
    Name      = "http_allow_all"
  }
}

resource "aws_lb_target_group" "product_service" {
  name_prefix = "prdtg"
  port        = 3010
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = data.aws_vpc.default.id
  lifecycle {
    create_before_destroy = true
  }
  tags = {
    Terraform = true
  }
}
