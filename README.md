# File processing

Processing a file when this is uploaded in a bucket

## Prerequisites
- Go >= 1.24
- AWS CLI
- Terraform
- LocalStack
- Docker (for LocalStack)

### Localstack
To use LocalStack on your local machine, add a profile in the aws cli settings: `.aws/credentials` and `.aws/config`

### Terraform State
The S3 bucket pre-built is used to store the terraform state file. The bucket name is defined in the `terraform-state.tf` file.


```shell
cd app/src
go build
go test
```

```shell
cd app/src

# Build
GOOS=linux GOARCH=amd64 go build -o ../target/file-processing main.go

# zip
zip ../target/app.zip ../target/file-processing
```

## Local Setup

1. Configure LocalStack profile:
```bash
aws configure set aws_access_key_id test --profile localstack
aws configure set aws_secret_access_key test --profile localstack
aws configure set region us-east-1 --profile localstack
aws configure set endpoint_url http://localhost:4566 --profile localstack
```

2. Create S3 bucket for Terraform state:
```bash
# Initialize localstack
localstack start -d

aws --endpoint-url=http://localhost:4566 s3 mb s3://alismed-terraform
```

## Terraform Commands
Local development

```bash
# Set AWS Profile for all commands
export AWS_PROFILE=localstack

# Initialize project:
terraform -chdir=infra init

# Validate configuration:
terraform -chdir=infra validate

# Format files:
terraform -chdir=infra fmt

# Plan changes:
terraform -chdir=infra plan

# Apply changes:
terraform -chdir=infra apply -auto-approve

# Destroy infrastructure:
terraform -chdir=infra destroy -auto-approve

# Unset AWS Profile (optional)
unset AWS_PROFILE

# Stop localstack
localstack stop
```