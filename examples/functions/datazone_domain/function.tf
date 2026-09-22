# arn:aws:datazone:ap-northeast-1:111111111111:domain/domain-id
output "datazone_domain" {
  value = provider::arn::datazone_domain("domain-id")
}
