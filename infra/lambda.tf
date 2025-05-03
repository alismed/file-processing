resource "aws_iam_role" "lambda_role" {
  name               = "file_processing_Role"
  assume_role_policy = file("iam/file_processing-role.json")
}

resource "aws_iam_policy" "iam_policy_for_lambda" {
  name        = "file_processing_policy"
  path        = "/"
  description = "AWS IAM Policy for managing aws lambda role"
  policy      = file("iam/file_processing-policy.json")
}

resource "aws_iam_role_policy_attachment" "attach_iam_policy_to_iam_role" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.iam_policy_for_lambda.arn
}

resource "aws_lambda_function" "file_processing" {
  function_name = "fileProcessing"
  filename      = "../app/target/app.zip"
  role          = aws_iam_role.lambda_role.arn
  handler       = "main"
  runtime       = "go1.x"
}