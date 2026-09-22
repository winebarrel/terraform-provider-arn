# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id
output "qbusiness_application" {
  value = provider::arn::qbusiness_application("application-id")
}
