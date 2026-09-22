# arn:aws:cloud9:ap-northeast-1:111111111111:environment:resource-id
output "cloud9_environment" {
  value = provider::arn::cloud9_environment("resource-id")
}
