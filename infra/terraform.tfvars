region          = "us-east-2"
profile         = "default"
bucket_name     = "alismed-file-to-upload"
storage_class   = "STANDARD"
table_name      = "orderform"
function_name   = "captureS3Upload"
event_rule_name = "catch-s3-uploads"
tags = {
  Environment = "production"
  Project     = "file processing"
}