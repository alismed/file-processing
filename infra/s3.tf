resource "aws_s3_bucket" "file_bucket" {
  bucket = var.bucket_name

  tags = merge(
    var.tags,
    {
      Name = "file-processing-bucket"
    }
  )
}
/*
resource "aws_s3_bucket_lifecycle_configuration" "file_bucket_lifecycle" {
  bucket = aws_s3_bucket.file_bucket.id

  rule {
    id     = "transition_to_ia"
    status = "Enabled"

    transition {
      days          = 30
      storage_class = "STANDARD_IA"
    }
  }
}

resource "aws_s3_object" "file_bucket_object" {
  bucket        = aws_s3_bucket.file_bucket.id
  key           = "file-processing-object"
  storage_class = "STANDARD_IA"

  tags = merge(
    var.tags,
    {
      Name = "file-processing-bucket"
    }
  )
}
*/