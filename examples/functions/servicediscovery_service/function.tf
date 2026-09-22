# arn:aws:servicediscovery:ap-northeast-1:111111111111:service/service-id
output "servicediscovery_service" {
  value = provider::arn::servicediscovery_service("service-id")
}
