resource "aws_cloudwatch_event_rule" "s3_upload" {
  name        = "capture-s3-uploads"
  description = "Capture all S3 object uploads"

  event_pattern = jsonencode({
    source      = ["aws.s3"]
    detail-type = ["Object Created"]
    detail = {
      bucket = {
        name = [aws_s3_bucket.file_bucket.id]
      }
    }
  })
}

resource "aws_cloudwatch_event_target" "s3_upload_target" {
  rule      = aws_cloudwatch_event_rule.s3_upload.name
  target_id = "ProcessS3Upload"
  arn       = aws_lambda_function.file_processing.arn
}