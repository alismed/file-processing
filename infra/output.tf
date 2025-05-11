output "aws_s3_bucket_name" {
  value = aws_s3_bucket.file_bucket.bucket
}

output "aws_s3_bucket_arn" {
  value = aws_s3_bucket.file_bucket.arn
}

output "aws_s3_bucket_id" {
  value = aws_s3_bucket.file_bucket.id
}

output "aws_s3_bucket_region" {
  value = aws_s3_bucket.file_bucket.region
}

output "aws_s3_bucket_domain_name" {
  value = aws_s3_bucket.file_bucket.bucket_domain_name
}

output "aws_lambda_function_arn" {
  value = aws_lambda_function.file_processing.arn
}

output "aws_lambda_function_name" {
  value = aws_lambda_function.file_processing.function_name
}

output "aws_lambda_function_role" {
  value = aws_iam_role.lambda_role.arn
}

output "aws_lambda_function_policy" {
  value = aws_iam_policy.iam_policy_for_lambda.arn
}

output "dynamodb_table_name" {
  value = aws_dynamodb_table.items.name
}

output "dynamodb_table_arn" {
  value = aws_dynamodb_table.items.arn
}

output "dynamodb_table_id" {
  value = aws_dynamodb_table.items.id
}
/*
output "aws_cloudwatch_event_rule_name" {
  value = aws_cloudwatch_event_rule.s3_upload.name
}

output "aws_cloudwatch_event_rule_arn" {
  value = aws_cloudwatch_event_rule.s3_upload.arn
}

output "aws_cloudwatch_event_rule_id" {
  value = aws_cloudwatch_event_rule.s3_upload.id
}
*/