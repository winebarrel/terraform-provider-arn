# arn:aws:dynamodb:ap-northeast-1:111111111111:table/table-name
output "dynamodb_table" {
  value = provider::arn::dynamodb_table("table-name")
}
