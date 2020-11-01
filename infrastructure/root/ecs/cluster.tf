resource "aws_ecs_cluster" "window_cluster" {
  name               = "window-cluster"
  capacity_providers = ["FARGATE"]
  tags = {
    Terraform = true
  }
}