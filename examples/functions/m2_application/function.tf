# arn:aws:m2:ap-northeast-1:111111111111:app/application-id
output "m2_application" {
  value = provider::arn::m2_application("application-id")
}
