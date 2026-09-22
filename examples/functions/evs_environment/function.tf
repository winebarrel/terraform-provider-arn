# arn:aws:evs:ap-northeast-1:111111111111:environment/environment-identifier
output "evs_environment" {
  value = provider::arn::evs_environment("environment-identifier")
}
