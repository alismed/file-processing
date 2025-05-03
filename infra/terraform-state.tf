terraform {
  backend "s3" {
    bucket  = "alismed-terraform2"
    key     = "file-processing/terraform.tfstate"
    region  = "us-east-2"
    encrypt = true
  }
}