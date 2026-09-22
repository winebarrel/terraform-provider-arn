# arn:aws:dynamodb::111111111111:global-table/global-table-name
output "dynamodb_global_table" {
  value = provider::arn::dynamodb_global_table("global-table-name")
}
