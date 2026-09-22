# arn:aws:cloudshell:ap-northeast-1:111111111111:environment/environment-id
output "cloudshell_environment" {
  value = provider::arn::cloudshell_environment("environment-id")
}
