# arn:aws:inspector2:ap-northeast-1:111111111111:connector/connector-id
output "inspector2_connector" {
  value = provider::arn::inspector2_connector("connector-id")
}
