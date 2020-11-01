resource "aws_ecr_repository" "go_rest_simple_product" {
  name                 = "go-rest-simple-product"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}