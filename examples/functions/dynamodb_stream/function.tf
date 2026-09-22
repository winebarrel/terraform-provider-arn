# arn:aws:dynamodb:ap-northeast-1:111111111111:table/table-name/stream/stream-label
output "dynamodb_stream" {
  value = provider::arn::dynamodb_stream("table-name", "stream-label")
}
