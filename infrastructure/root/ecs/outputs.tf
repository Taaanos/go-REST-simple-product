output "product_service_revision" {
  value = aws_ecs_task_definition.product_service.revision
}

output "lb_dns" {
  value = aws_lb.window_lb.dns_name
}