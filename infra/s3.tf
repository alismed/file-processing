resource "aws_s3_bucket" "file_bucket" {
  bucket = var.bucket_name

  tags = merge(
    var.tags,
    {
      Name = "file-processing-bucket"
    }
  )
}