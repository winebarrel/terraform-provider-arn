# arn:aws:m2:ap-northeast-1:111111111111:env/environment-id
output "m2_environment" {
  value = provider::arn::m2_environment("environment-id")
}
