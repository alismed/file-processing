/*

# Terraform backend configuration for localstack
# This configuration is used to store the terraform state file in a localstack s3 bucket
# This configuration is used only for local development

terraform {
  backend "s3" {
    bucket                      = "alismed-terraform"
    key                         = "file-processing/terraform.tfstate"
    region                      = "us-east-1"
    endpoint                    = "http://localhost:4566"
    skip_credentials_validation = true
    skip_metadata_api_check     = true
    force_path_style           = true
    profile                    = "localstack"
  }
}
*/