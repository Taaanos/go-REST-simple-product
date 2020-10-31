resource "aws_ecs_cluster" "window_cluster" {
  name               = "window-cluster"
  capacity_providers = ["FARGATE"]
  tags = {
    Terraform = true
  }
}

resource "aws_ecs_task_definition" "product_service" {
  family                   = "product-service"
  container_definitions    = file("task-definitions/product-service.json")
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 256
  memory                   = 512
  execution_role_arn       = aws_iam_role.ecs_task_role.arn
  tags = {
    Terraform = true
  }
}

resource "aws_ecs_service" "product_service" {
  name            = "product-service"
  cluster         = aws_ecs_cluster.window_cluster.id
  task_definition = aws_ecs_task_definition.product_service.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  network_configuration {
    subnets          = data.aws_subnet_ids.default.ids
    security_groups  = [aws_security_group.http_allow_all.id]
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

output "product_service_revision" {
  value = aws_ecs_task_definition.product_service.revision
}


resource "aws_iam_role_policy" "ecs_task_policy" {
  name = "ecs-task-policy"
  role = aws_iam_role.ecs_task_role.id

  policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ecr:GetAuthorizationToken",
                "ecr:BatchCheckLayerAvailability",
                "ecr:GetDownloadUrlForLayer",
                "ecr:BatchGetImage",
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": "*"
        }
    ]
  }
  EOF
}

resource "aws_iam_role" "ecs_task_role" {
  name = "ecs-task-role"

  assume_role_policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Principal": {
          "Service": "ecs-tasks.amazonaws.com"
        },
        "Effect": "Allow",
        "Sid": ""
      }
    ]
  }
  EOF
}