# arn:aws:batch:ap-northeast-1:111111111111:service-environment/service-environment-name
output "batch_service_environment" {
  value = provider::arn::batch_service_environment("service-environment-name")
}
