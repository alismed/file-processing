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

variable "table_name" {
  description = "The name of the DynamoDB table"
  type        = string
}

variable "billing_mode" {
  description = "The billing mode for the DynamoDB table"
  type        = string
  default     = "PAY_PER_REQUEST"
}

variable "ttl_attribute" {
  description = "The attribute name for TTL in the DynamoDB table"
  type        = string
  default     = "ttl"
}

variable "gsi_config" {
  description = "Configuration for global secondary indexes"
  type = map(object({
    hash_key           = string
    write_capacity     = number
    read_capacity      = number
    projection_type    = string
    non_key_attributes = list(string)
  }))
  default = {}
}


