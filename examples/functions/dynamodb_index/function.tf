# arn:aws:dynamodb:ap-northeast-1:111111111111:table/table-name/index/index-name
output "dynamodb_index" {
  value = provider::arn::dynamodb_index("table-name", "index-name")
}
