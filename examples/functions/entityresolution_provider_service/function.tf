# arn:aws:entityresolution:ap-northeast-1:111111111111:providerservice/provider-name/provider-service-name
output "entityresolution_provider_service" {
  value = provider::arn::entityresolution_provider_service("provider-name", "provider-service-name")
}
