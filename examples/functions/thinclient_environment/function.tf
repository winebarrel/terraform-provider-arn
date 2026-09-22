# arn:aws:thinclient:ap-northeast-1:111111111111:environment/environment-id
output "thinclient_environment" {
  value = provider::arn::thinclient_environment("environment-id")
}
