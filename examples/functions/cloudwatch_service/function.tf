# arn:aws:cloudwatch:ap-northeast-1:111111111111:service/service-name-unique-attributes-hex
output "cloudwatch_service" {
  value = provider::arn::cloudwatch_service("service-name", "unique-attributes-hex")
}
