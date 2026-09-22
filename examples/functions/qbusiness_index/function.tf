# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id/index/index-id
output "qbusiness_index" {
  value = provider::arn::qbusiness_index("application-id", "index-id")
}
