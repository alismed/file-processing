variable "region" {
  description = "The AWS region to deploy the infrastructure"
  type        = string
}

variable "profile" {
  description = "The AWS profile to use for authentication"
  type        = string
}

variable "tags" {
  description = "Tags to apply to all resources"
  type        = map(string)
  default = {
    Environment = "production"
    Project     = "file processing"
  }
}

variable "bucket_name" {
  description = "The name of the S3 bucket to store the Terraform state file"
  type        = string
}

variable "storage_class" {
  description = "The storage class for the S3 bucket"
  type        = string
  default     = "STANDARD_IA"
}


