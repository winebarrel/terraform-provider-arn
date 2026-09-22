# arn:aws:ssm:ap-northeast-1:111111111111:servicesetting/resource-id
output "ssm_servicesetting" {
  value = provider::arn::ssm_servicesetting("resource-id")
}
