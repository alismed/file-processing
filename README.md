# File processing

Processing a file when this is uploaded in a bucket, creating new records on a database

## Prerequisites
- Go >= 1.24
- AWS CLI
- Terraform
- LocalStack
- Docker (for LocalStack)

### Localstack
To use LocalStack on your local machine, add a profile in the aws cli settings: `.aws/credentials` and `.aws/config`

### Terraform State
The S3 bucket pre-built is used to store the terraform state file.

If using localstack the bucket name is defined in the `terraform-state.local.tf` file. Else in the `terraform-state.tf`. Only one should be active.

### Build and test local the lambda code
```shell
cd app/src

# Build
GOOS=linux GOARCH=amd64 go build -o ../target/file-processing main.go

# Test
go test

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

## GitHub Actions Workflow

The deployment process is automated using GitHub Actions with three stages:

1. **Validation**
   - Terraform validation
   - Variable checking
2. **Terraform Execution**
   - AWS credentials setup
   - Infrastructure deployment
3. **PR Notification**
   - Deployment status updates
   - PR comments


### Testing Actions Locally

1. Install Act:
```bash
curl -s https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
```

2. Setup test environment:
```bash
# Create test directory if not exists
mkdir -p .act

# Create env file with credentials
echo "AWS_ACCESS_KEY_ID=test" > .act/.env
echo "AWS_SECRET_ACCESS_KEY=test" >> .act/.env
echo "AWS_DEFAULT_REGION=us-east-1" >> .act/.env

# Create pull request event simulation
cat > .act/pull_request.json << EOF
{
  "pull_request": {
    "number": 1,
    "body": "Test PR",
    "head": {
      "ref": "feature/test"
    }
  }
}
EOF
```

3. Run test workflow:
```bash
# List available workflows
act -l

# Run workflow with pull request event
act pull_request -e .act/pull_request.json --secret-file .act/.env

# Run specific workflow
act -W .github/workflows/main.yaml \
    -e .act/pull_request.json \
    --secret-file .act/.env \
    --container-architecture linux/amd64

# Run with verbose output
act -v pull_request -e .act/pull_request.json --secret-file .act/.env
```