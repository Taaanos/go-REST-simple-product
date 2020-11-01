terraform {
  backend "s3" {
    bucket         = "terraform-state20201101155213607300000001"
    key            = "backends/terraform.tfstate"
    region         = "eu-west-2"
    encrypt        = true
    dynamodb_table = "state-lock"
  }
}

resource "aws_s3_bucket" "terraform_state" {
  bucket_prefix = "terraform-state"
  versioning {
    enabled = true
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_dynamodb_table" "state_lock" {
  name         = "state-lock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}

output "bucket_name" {
  value = aws_s3_bucket.terraform_state.id
}