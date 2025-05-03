resource "aws_dynamodb_table" "items" {
  name         = var.table_name
  billing_mode = var.billing_mode
  hash_key     = "Id"

  point_in_time_recovery {
    enabled = false
  }

  ttl {
    enabled        = true
    attribute_name = var.ttl_attribute
  }

  attribute {
    name = "Id"
    type = "S"
  }

  attribute {
    name = "name"
    type = "S"
  }

  attribute {
    name = "region"
    type = "N"
  }

  attribute {
    name = "description"
    type = "S"
  }

  global_secondary_index {
    name            = "NameIndex"
    hash_key        = "name"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "RegionIndex"
    hash_key        = "region"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "DescriptionIndex"
    hash_key        = "description"
    projection_type = "ALL"
  }

  tags = merge(
    var.tags,
    {
      Name = "file-processing-database"
    }
  )
}