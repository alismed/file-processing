/*
resource "aws_cloudwatch_event_rule" "s3_upload" {
  name        = var.event_rule_name
  description = "Capture all S3 object uploads"

  event_pattern = jsonencode({
    source      = ["aws.s3"]
    detail-type = ["AWS API Call via CloudTrail"]
    detail = {
      eventSource = ["s3.amazonaws.com"]
      eventName   = ["PutObject", "CompleteMultipartUpload"]
      requestParameters = {
        bucketName = [aws_s3_bucket.file_bucket.id]
      }
    }
  })
}

resource "aws_cloudwatch_event_target" "s3_upload_target" {
  rule      = aws_cloudwatch_event_rule.s3_upload.name
  target_id = var.event_target_id
  arn       = aws_lambda_function.file_processing.arn
}
*/