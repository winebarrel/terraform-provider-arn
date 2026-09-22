# arn:aws:apprunner:ap-northeast-1:111111111111:service/service-name/service-id
output "apprunner_service" {
  value = provider::arn::apprunner_service("service-name", "service-id")
}
