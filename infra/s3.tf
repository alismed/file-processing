resource "aws_s3_bucket" "file_bucket" {
  bucket = var.bucket_name

  tags = merge(
    var.tags,
    {
      Name = "file-processing-bucket"
    }
  )
}

resource "aws_lambda_permission" "allow_s3_invoke" {
  #depends_on = [aws_s3_bucket.file_bucket]
  statement_id  = "AllowExecutionFromS3"
  action        = "lambda:InvokeFunction"
  function_name = var.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.file_bucket.arn
}
/*
resource "aws_lambda_permission" "allow_eventbridge" {
  statement_id  = "AllowEventBridgeInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.file_processing.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.s3_upload.arn
}
*/
resource "aws_s3_bucket_notification" "bucket_notification" {
  depends_on = [aws_lambda_permission.allow_s3_invoke]
  bucket     = aws_s3_bucket.file_bucket.id
  lambda_function {
    lambda_function_arn = aws_lambda_function.file_processing.arn
    events              = ["s3:ObjectCreated:*"]
    #filter_prefix       = "uploads/"
    #filter_suffix       = ".csv"
  }

  # eventbridge = true
}
