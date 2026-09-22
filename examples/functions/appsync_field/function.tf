# arn:aws:appsync:ap-northeast-1:111111111111:apis/graph-qlapi-id/types/type-name/fields/field-name
output "appsync_field" {
  value = provider::arn::appsync_field("graph-qlapi-id", "type-name", "field-name")
}
