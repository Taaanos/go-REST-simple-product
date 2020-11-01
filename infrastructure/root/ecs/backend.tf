terraform {
  backend "s3" {
    bucket         = "terraform-state20201101155213607300000001"
    key            = "ecs/terraform.tfstate"
    region         = "eu-west-2"
    encrypt        = true
    dynamodb_table = "state-lock"
  }
}