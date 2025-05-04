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

resource "null_resource" "go_build" {
  triggers = {
    always_run = "${timestamp()}" # Ensure the build executes each apply
  }

  provisioner "local-exec" {
    command = <<EOT
      mkdir -p ../app/target
      cd ../app/src
      GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../target/bootstrap main.go
    EOT
  }
}

data "archive_file" "lambda_app" {
  depends_on  = [null_resource.go_build]
  type        = "zip"
  source_file = "../app/target/bootstrap"
  output_path = "../app/target/package/app.zip"
}

resource "aws_cloudwatch_log_group" "lambda_log_group" {
  name              = "/aws/lambda/${var.function_name}"
  retention_in_days = 7
  tags = merge(
    var.tags,
    {
      Name = "file-processing-log-group"
    }
  )
}

resource "aws_lambda_function" "file_processing" {
  filename         = data.archive_file.lambda_app.output_path
  function_name    = var.function_name
  role             = aws_iam_role.lambda_role.arn
  handler          = "bootstrap"
  runtime          = "provided.al2"
  source_code_hash = data.archive_file.lambda_app.output_base64sha256

  depends_on = [
    aws_iam_role_policy_attachment.attach_iam_policy_to_iam_role,
    data.archive_file.lambda_app,
    aws_cloudwatch_log_group.lambda_log_group
  ]
}
