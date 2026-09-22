# arn:aws:servicequotas:ap-northeast-1:111111111111:service-code/quota-code
output "servicequotas_quota" {
  value = provider::arn::servicequotas_quota("service-code", "quota-code")
}
