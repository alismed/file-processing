resource "aws_s3_bucket" "file_bucket" {
  bucket = var.bucket_name

  tags = merge(
    var.tags,
    {
      Name = "file-processing-bucket"
    }
  )
}

resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.file_bucket.id

  eventbridge = true
}
