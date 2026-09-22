# arn:aws:aidevops:ap-northeast-1:111111111111:service/service-id
output "aidevops_service" {
  value = provider::arn::aidevops_service("service-id")
}
