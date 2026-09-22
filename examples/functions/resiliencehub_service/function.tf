# arn:aws:resiliencehub:ap-northeast-1:111111111111:service/service-id
output "resiliencehub_service" {
  value = provider::arn::resiliencehub_service("service-id")
}
